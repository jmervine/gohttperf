package gohttperf

import (
    "testing"
)

func TestCalculatePercentiles(t *testing.T) {
    var times []float64
    for i := 1.1; i < 101; i++ {
        times = append(times, float64(i))
    }

    result := Results{ConnectionTimes: times}
    result.calculatePercentiles()

    var pct float64
    pct = result.ConnectionTime75Pct
    if pct != 75.1 {
        t.Error("Expected '75.1', but got", pct)
    }

    pct = result.ConnectionTime80Pct
    if pct != 80.1 {
        t.Error("Expected '80.1', but got", pct)
    }

    pct = result.ConnectionTime85Pct
    if pct != 85.1 {
        t.Error("Expected '85.1', but got", pct)
    }

    pct = result.ConnectionTime90Pct
    if pct != 90.1 {
        t.Error("Expected '90.1', but got", pct)
    }

    pct = result.ConnectionTime95Pct
    if pct != 95.1 {
        t.Error("Expected '95.1', but got", pct)
    }

    pct = result.ConnectionTime99Pct
    if pct != 99.1 {
        t.Error("Expected '99.1', but got", pct)
    }
}

func TestPercentiles(t *testing.T) {
    var times []float64
    for i := 1.1; i < 101; i++ {
        times = append(times, float64(i))
    }

    var pct float64
    pct = percentile(times, 75)
    if pct != 75.1 {
        t.Error("Expected '75.1', but got", pct)
    }

    pct = percentile(times, 80)
    if pct != 80.1 {
        t.Error("Expected '80.1', but got", pct)
    }

    pct = percentile(times, 85)
    if pct != 85.1 {
        t.Error("Expected '85.1', but got", pct)
    }

    pct = percentile(times, 90)
    if pct != 90.1 {
        t.Error("Expected '90.1', but got", pct)
    }

    pct = percentile(times, 95)
    if pct != 95.1 {
        t.Error("Expected '95.1', but got", pct)
    }

    pct = percentile(times, 99)
    if pct != 99.1 {
        t.Error("Expected '99.1', but got", pct)
    }
}

func TestFind(t *testing.T) {
    result := find("foobar\nbahboo\nack\n", "^(bah.+)$")

    if result != "bahboo" {
        t.Error("Expected 'bahboo', but got", result)
    }
}

func TestFindConnectionTimes(t *testing.T) {
    raw := "Connection lifetime = 1.0\nConnection lifetime = 2.0\nConnection lifetime = 3.0\n"

    result := findConnectionTimes(raw)

    if len(result) != 3 {
        t.Error("Expected '3', but got", len(result))
    }

    if result[1] != 2.0 {
        t.Error("Expected '2.0', but got", result[1])
    }
}

func TestToS(t *testing.T) {
    if toS("foo") != "foo" {
        t.Error("Expected 'foo', but got", toS("foo"))
    }
}

func TestToI(t *testing.T) {
    i := toI("10")
    if i != 10 {
        t.Error("Expected '10', but got", i)
    }
}

func TestToF(t *testing.T) {
    f := toF("10.1")
    if f != 10.1 {
        t.Error("Expected '10.1', but got", f)
    }
}
