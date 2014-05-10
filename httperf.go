/*

Package gohttperf is a simple wrapper for httperf.

Basic Run Example:

    // Define httperf arguments
    options := map[string]interface{}{
        "server": "localhost",
        "uri":    "/foo",
        "hog":    true,
    }

    // Create HTTPerf
    httperf := HTTPerf{ Options: options }

    // Run httperf, catching any errors with return value.
    if err := httperf.Run(); err == nil {
        fmt.Print(httperf.Raw)
    }

    // Turn on Parser
    httperf.Parser = true

    // Run httperf, catching any errors with return value.
    if err := httperf.Run(); err == nil {
        fmt.Println("QPS: ", httperf.Results.ConnectionRatePerSec)
        fmt.Println("200s:", httperf.Results.ReplyStatus2xx)
    }


Basic Fork Example:

    var output bytes.Buffer

    cmd, err := httperf.Fork(&output)
    deferred := func() {
        if err := httperf.Wait(cmd, &output); err == nil {
            fmt.Println("Parsed:")
            fmt.Println("-------")
            fmt.Println("QPS: ", httperf.Results.ConnectionRatePerSec)
            fmt.Println("200s:", httperf.Results.ReplyStatus2xx)
        }
    }

    if err == nil {
        defer deferred()
    }

    // do something before calling wait
    fmt.Println("I'm waiting...\n")

*/
package gohttperf

import (
    "bytes"
    "github.com/jmervine/sh"
    "os/exec"
    "strconv"
    "strings"
)

// HTTPerf is the main data struct, (almost) all things use h.
type HTTPerf struct {
    Options map[string]interface{}
    Path    string
    Parser  bool
    Raw     string
    Results Results
}

// Turns (*HTTPerf).Options to a string of command line arguments.
//
//   When:
//   (*HTTPerf).Options = map[string]interface{}{
//      "hog": true,
//      "verbose": true,
//      "server": "localhost"
//   }
//
//  Returns:
//
//  "--hog --verbose --server localhost"
func (h *HTTPerf) arguments() string {
    var params = map[string]interface{}{
        "hog":              false,
        "verbose":          false,
        "add-header":       nil,
        "burst-length":     nil,
        "client":           nil,
        "close-with-reset": nil,
        "debug":            nil,
        "failure-status":   nil,
        "http-version":     nil,
        "max-connections":  nil,
        "max-piped-calls":  nil,
        "method":           nil,
        "no-host-hdr":      nil,
        "num-calls":        nil,
        "num-conns":        nil,
        "period":           nil,
        "port":             nil,
        "print-reply":      nil,
        "print-request":    nil,
        "rate":             nil,
        "recv-buffer":      nil,
        "retry-on-failure": nil,
        "send-buffer":      nil,
        "server":           nil,
        "server-name":      nil,
        "session-cookies":  nil,
        "ssl":              nil,
        "ssl-ciphers":      nil,
        "ssl-no-reuse":     nil,
        "think-timeout":    nil,
        "timeout":          nil,
        "uri":              nil,
        "wlog":             nil,
        "wsess":            nil,
        "wsesslog":         nil,
        "wset":             nil,
    }

    var args []string
    for key, val := range h.Options {
        if _, ok := params[key]; ok {
            params[key] = val
        }
    }

    for key, val := range params {
        var arg string

        if val == nil || val == false {
            continue
        }

        if key == "uri" {
            args = append(args, "--uri '"+val.(string)+"'")
            continue
        }

        switch val.(type) {
        case bool:
            if val.(bool) {
                arg = "--" + key
            }
        case int:
            arg = "--" + key + " " + strconv.Itoa(val.(int))
        case string:
            arg = "--" + key + " " + val.(string)
        }
        args = append(args, arg)
    }

    return strings.Join(args, " ")
}

// Command returns (*HTTPerf).Path + (*HTTPerf).arguments()
//
//   When:
//   (*HTTPerf).Path = "httperf" // the default
//   (*HTTPerf).Options = map[string]interface{}{
//      "hog": true,
//      "verbose": true,
//      "server": "localhost"
//   }
//
//  Returns:
//
//  "httperf --hog --verbose --server localhost"
func (h *HTTPerf) Command() string {
    cmd := ""

    if h.Path == "" {
        cmd += "httperf"
    } else {
        cmd += h.Path
    }

    cmd += " "
    cmd += h.arguments()
    return cmd
}

// Run executes the command string returned by (*HTTPerf).Command()
// on the shell.
func (h *HTTPerf) Run() error {
    result, err := sh.Sh(h.Command())
    h.Raw = string(result[:])

    if h.Parser {
        // Sets h.Results to parsed Results struct.
        h.Parse()
    }

    return err
}

// Fork executes the command string returned by (*HTTPerf).Command()
// on the shell in a async manner, returning the *exec.Cmd pointer
// to be referenced in Wait when ready to gather results.
func (h *HTTPerf) Fork(output *bytes.Buffer) (*exec.Cmd, error) {
    return sh.ShFork(h.Command(), output)
}

// Wait takes the *exec.Cmd returned by Fork and waits for it to complete
// it's run, on completion it sets (*HTTPerf).Raw and runs the parser
// if requested.
func (h *HTTPerf) Wait(cmd *exec.Cmd, output *bytes.Buffer) error {
    result, err := sh.ShWait(cmd, output)

    if err == nil {
        h.Raw = result

        if h.Parser {
            // Sets h.Results to parsed Results struct.
            h.Parse()
        }
    }

    return err
}
