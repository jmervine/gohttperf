// file: utils.go - Non-exported utility functions used throughout.

package gohttperf

import (
    "errors"
    "math"
    "regexp"
    "sort"
    "strconv"
    "strings"
)

// When an existing Results struct contains ConnectionTimes (which
// will only be the case on a verbose httperf run using the fork of
// httperf found at https://github.com/jmervine/httperf) this method
// will be run to calculate the 75th, 80th, 85th, 90th, 95th and 99th
// percentiles of the run.
func (r *Results) calculatePercentiles() {
    times := r.ConnectionTimes
    if len(times) > 0 {
        r.ConnectionTime75Pct = percentile(times, 75)
        r.ConnectionTime80Pct = percentile(times, 80)
        r.ConnectionTime85Pct = percentile(times, 85)
        r.ConnectionTime90Pct = percentile(times, 90)
        r.ConnectionTime95Pct = percentile(times, 95)
        r.ConnectionTime99Pct = percentile(times, 99)
    }
}

// Helper method for (*Results).calculatePercentiles which actually
// finds the percentile from the passed slice of connection times.
func percentile(times []float64, pct int) float64 {
    length := len(times)

    if length == 1 {
        return times[0]
    }

    if length == 2 {
        return times[1]
    }

    var times64 []float64
    for _, float := range times {
        times64 = append(times64, float64(float))
    }

    if !sort.Float64sAreSorted(times64) {
        sort.Float64s(times64)
    }

    index := int(math.Floor(((float64(length)/100)*float64(pct))+0.5)) - 1

    return float64(times64[index])
}

// Helper method which wraps and simplifies regexp methods
// to easily turn a needle (pattern string) and haystack
// (raw string) in to a result string.
func find(raw string, pattern string) string {
    expression := regexp.MustCompile(pattern)
    lines := strings.Split(raw, "\n")

    for _, line := range lines {
        if line != "" {
            match := expression.FindStringSubmatch(line)
            if len(match) == 2 {
                return match[1]
            }
        }
    }
    return ""
}

// Helper method to search raw results for 'Connection lifetime' output
// and return all results as a []float64.
func findConnectionTimes(raw string) []float64 {
    var result []float64
    lines := strings.Split(raw, "\n")

    verboseExpression := regexp.MustCompile("^Connection lifetime = ([0-9]+\\.[0-9]+)(\\s?)")
    for _, line := range lines {
        if verboseExpression.MatchString(line) {
            result = append(result, toF(verboseExpression.FindStringSubmatch(line)[1]))
        }
    }
    return result
}

// Error handling around strings, for use with regexp parsing
// raw results. These results should never be empty.
func toS(s string) string {
    if s == "" {
        panic(errors.New(s + " should not be empty"))
    }
    return s
}

// Convert string to int and handle any errors, for use with regexp
// parsing raw results. These results should never be empty.
func toI(s string) int {
    if s == "" {
        panic(errors.New(s + " should not be empty"))
    }
    i, e := strconv.ParseInt(s, 0, 16)
    if e != nil {
        panic(e)
    }
    return int(i)
}

// Convert string to float and handle any errors, for use with regexp
// parsing raw results. These results should never be empty.
func toF(s string) float64 {
    if s == "" {
        panic(errors.New(s + " should not be empty"))
    }

    f, e := strconv.ParseFloat(s, 64)
    if e != nil {
        panic(e)
    }
    return f
}
