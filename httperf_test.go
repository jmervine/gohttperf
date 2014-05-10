package gohttperf

import (
    "bytes"
    "fmt"
    "os"
    "regexp"
    "testing"
)

func TestInitWithPath(t *testing.T) {
    httperf := HTTPerf{}
    if httperf.Path != "" {
        t.Error("Expected '', got ", httperf.Path)
    }

    if httperf.Command() != "httperf " {
        t.Error("Expected 'httperf ', got '" + httperf.Command() + "'")
    }

    httperf = HTTPerf{Path: "foobar"}

    if httperf.Path != "foobar" {
        t.Error("Expected 'foobar', got ", httperf.Path)
    }

    if httperf.Command() != "foobar " {
        t.Error("Expected 'foobar ', got '" + httperf.Command() + "'")
    }
}

func TestInitWithParser(t *testing.T) {
    httperf := HTTPerf{}

    if httperf.Parser {
        t.Error("Expected parse == false")
    }

    httperf.Parser = true
    if !httperf.Parser {
        t.Error("Expected parse == true")
    }
}

func TestInitWithOptions(t *testing.T) {
    httperf := HTTPerf{
        Options: map[string]interface{}{
            "server":  "www.example.com",
            "verbose": true,
        },
    }

    if httperf.Options["server"] != "www.example.com" {
        t.Error("Expected server == 'www.example.com', got ", httperf.Options["server"])
    }

    if !httperf.Options["verbose"].(bool) {
        t.Error("Expected verbose == true, got ", httperf.Options["verbose"])
    }

    var (
        re  string
        cmd string
        ok  bool
        err error
    )

    re = "--server www.example.com"
    cmd = httperf.Command()
    if ok, err = regexp.MatchString(re, cmd); !ok || err != nil {
        t.Error("Expected '" + re + "' in '" + cmd + "'")
    }

    re = "--verbose"
    cmd = httperf.Command()
    if ok, err = regexp.MatchString(re, cmd); !ok || err != nil {
        t.Error("Expected '" + re + "' in '" + cmd + "'")
    }

    httperf.Options["hog"] = true
    httperf.Options["rate"] = 10

    re = "--server www.example.com"
    cmd = httperf.Command()
    if ok, err = regexp.MatchString(re, cmd); !ok || err != nil {
        t.Error("Expected '" + re + "' in '" + cmd + "'")
    }

    re = "--verbose"
    cmd = httperf.Command()
    if ok, err = regexp.MatchString(re, cmd); !ok || err != nil {
        t.Error("Expected '" + re + "' in '" + cmd + "'")
    }

    re = "--hog"
    cmd = httperf.Command()
    if ok, err = regexp.MatchString(re, cmd); !ok || err != nil {
        t.Error("Expected '" + re + "' in '" + cmd + "'")
    }

    re = "--rate 10"
    cmd = httperf.Command()
    if ok, err = regexp.MatchString(re, cmd); !ok || err != nil {
        t.Error("Expected '" + re + "' in '" + cmd + "'")
    }
}

func TestInitWithBadOption(t *testing.T) {
    httperf := HTTPerf{
        Options: map[string]interface{}{
            "bad": "bad",
        },
    }

    if httperf.Command() != "httperf " {
        t.Error("Expected 'httperf ', got '" + httperf.Command() + "'")
    }

}

func TestArguments(t *testing.T) {
    httperf := HTTPerf{
        Options: map[string]interface{}{
            "server":  "localhost",
            "uri":     "/foo",
            "hog":     true,
            "verbose": true,
        },
    }

    var (
        re   string
        args string
        ok   bool
        err  error
    )

    args = httperf.arguments()
    re = "--hog"
    if ok, err = regexp.MatchString(re, args); !ok || err != nil {
        t.Error("Expected '" + re + "' in '" + args + "'")
    }

    re = "--server localhost"
    if ok, err = regexp.MatchString(re, args); !ok || err != nil {
        t.Error("Expected '" + re + "' in '" + args + "'")
    }

    re = "--uri '/foo'"
    if ok, err = regexp.MatchString(re, args); !ok || err != nil {
        t.Error("Expected '" + re + "' in '" + args + "'")
    }

    re = "--verbose"
    if ok, err = regexp.MatchString(re, args); !ok || err != nil {
        t.Error("Expected '" + re + "' in '" + args + "'")
    }

}

func TestRun(t *testing.T) {
    httperf := HTTPerf{
        Options: map[string]interface{}{
            "server": "localhost",
            "uri":    "/foo",
            "hog":    true,
        },
    }

    root, _ := os.Getwd()
    httperf.Path = root + "/_support/httperf"

    err := httperf.Run()

    if err != nil {
        t.Error("Expected no error.")
    }

    match, err := regexp.MatchString("^httperf ", httperf.Raw)

    if err != nil {
        t.Error("Expected no error.")
    }

    if !match {
        t.Error("Expected match.")
    }
}

func TestForkAndWait(t *testing.T) {
    httperf := HTTPerf{
        Options: map[string]interface{}{
            "server": "localhost",
            "uri":    "/foo",
            "hog":    true,
        },
    }

    root, _ := os.Getwd()
    httperf.Path = root + "/_support/httperf"

    var output bytes.Buffer

    cmd, err := httperf.Fork(&output)

    if err != nil {
        t.Error("Expected no error.")
    }

    err = httperf.Wait(cmd, &output)

    if err != nil {
        t.Error("Expected no error.")
    }

    match, err := regexp.MatchString("^httperf ", httperf.Raw)

    if err != nil {
        t.Error("Expected no error.")
    }

    if !match {
        t.Error("Expected match.")
    }
}

func TestRunWithParser(t *testing.T) {
    httperf := HTTPerf{
        Options: map[string]interface{}{
            "server": "localhost",
            "uri":    "/foo",
            "hog":    true,
        },
        Parser: true,
    }

    root, _ := os.Getwd()
    httperf.Path = root + "/_support/httperf"

    err := httperf.Run()

    if err != nil {
        t.Error("Expected no error.")
    }

    match, err := regexp.MatchString("^httperf ", httperf.Raw)

    if err != nil {
        t.Error("Expected no error.")
    }

    if !match {
        t.Error("Expected match.")
    }

    if httperf.Results.TotalConnections != 1 {
        t.Error("Expected 1, but got", httperf.Results.ConnectionRatePerSec)
    }

    if httperf.Results.ConnectionRatePerSec != 4524.6 {
        t.Error("Expected 4524.6, but got", httperf.Results.ConnectionRatePerSec)
    }
}

func TestForkAndWaitWithParser(t *testing.T) {
    httperf := HTTPerf{
        Options: map[string]interface{}{
            "server": "localhost",
            "uri":    "/foo",
            "hog":    true,
        },
        Parser: true,
    }

    root, _ := os.Getwd()
    httperf.Path = root + "/_support/httperf"

    var output bytes.Buffer

    cmd, err := httperf.Fork(&output)

    if err != nil {
        t.Error("Expected no error.")
    }

    err = httperf.Wait(cmd, &output)

    if err != nil {
        t.Error("Expected no error.")
    }

    match, err := regexp.MatchString("^httperf ", httperf.Raw)

    if err != nil {
        t.Error("Expected no error.")
    }

    if !match {
        t.Error("Expected match.")
    }

    if httperf.Results.TotalConnections != 1 {
        t.Error("Expected 1, but got", httperf.Results.ConnectionRatePerSec)
    }

    if httperf.Results.ConnectionRatePerSec != 4524.6 {
        t.Error("Expected 4524.6, but got", httperf.Results.ConnectionRatePerSec)
    }
}

func TestRunVerbose(t *testing.T) {
    httperf := HTTPerf{
        Options: map[string]interface{}{
            "server":  "localhost",
            "uri":     "/foo",
            "hog":     true,
            "verbose": true,
        },
    }

    root, _ := os.Getwd()
    httperf.Path = root + "/_support/httperf"

    err := httperf.Run()

    if err != nil {
        t.Error("Expected no error.")
    }

    match, err := regexp.MatchString("^httperf ", httperf.Raw)

    if err != nil {
        t.Error("Expected no error.")
    }

    if !match {
        t.Error("Expected match.")
    }
}

func TestRunVerboseWithParser(t *testing.T) {
    httperf := HTTPerf{
        Options: map[string]interface{}{
            "server":  "localhost",
            "uri":     "/foo",
            "hog":     true,
            "verbose": true,
        },
        Parser: true,
    }

    root, _ := os.Getwd()
    httperf.Path = root + "/_support/httperf"

    err := httperf.Run()

    if err != nil {
        t.Error("Expected no error.")
    }

    match, err := regexp.MatchString("^httperf ", httperf.Raw)

    if err != nil {
        t.Error("Expected no error.")
    }

    if !match {
        t.Error("Expected match.")
    }

    if httperf.Results.TotalConnections != 100 {
        t.Error("Expected 100, but got", httperf.Results.ConnectionRatePerSec)
    }

    if httperf.Results.ConnectionRatePerSec != 4.0 {
        t.Error("Expected 4.0, but got", httperf.Results.ConnectionRatePerSec)
    }

    //if httperf.Results.ConnectionTime75Pct != 0.0 {
    //t.Error("Expected 0.0, but got", httperf.Results.ConnectionTime75Pct)
    //}
}

func ExampleHTTPerf_Run() {
    /* Define httperf arguments */
    options := map[string]interface{}{
        "server": "localhost",
        "uri":    "/foo",
        "hog":    true,
    }

    /* Create HTTPerf */
    httperf := HTTPerf{
        /* Stub path for testing. */
        Path:    "./_support/httperf",
        Options: options,
    }

    fmt.Println("Raw:")
    fmt.Println("-------")

    /* Run httperf, catching any errors with return value. */
    if err := httperf.Run(); err == nil {
        /* Print raw results (stdout and stderr) if successful. */
        fmt.Print(httperf.Raw)
    }

    // Output:
    //Raw:
    //-------
    //httperf --server localhost --uri /foo --hog
    //httperf: warning: open file limit > FD_SETSIZE; limiting max. # of open files to FD_SETSIZE
    //Maximum connect burst length: 0
    //
    //Total: connections 1 requests 1 replies 1 test-duration 0.000 s
    //
    //Connection rate: 4524.6 conn/s (0.2 ms/conn, <=1 concurrent connections)
    //Connection time [ms]: min 0.2 avg 0.2 max 0.2 median 0.5 stddev 0.0
    //Connection time [ms]: connect 0.1
    //Connection length [replies/conn]: 1.000
    //
    //Request rate: 4524.6 req/s (0.2 ms/req)
    //Request size [B]: 62.0
    //
    //Reply rate [replies/s]: min 0.0 avg 0.0 max 0.0 stddev 0.0 (0 samples)
    //Reply time [ms]: response 0.1 transfer 0.0
    //Reply size [B]: header 154.0 content 168.0 footer 0.0 (total 322.0)
    //Reply status: 1xx=0 2xx=0 3xx=0 4xx=1 5xx=0
    //
    //CPU time [s]: user 0.00 system 0.00 (user 0.0% system 0.0% total 0.0%)
    //Net I/O: 1696.7 KB/s (13.9*10^6 bps)
    //
    //Errors: total 0 client-timo 0 socket-timo 0 connrefused 0 connreset 0
    //Errors: fd-unavail 0 addrunavail 0 ftab-full 0 other 0
}

func ExampleHTTPerf_Fork() {
    /* Define httperf arguments */
    options := map[string]interface{}{
        "server": "localhost",
        "uri":    "/foo",
        "hog":    true,
    }

    /* Create HTTPerf */
    httperf := HTTPerf{
        /* Stub path for testing. */
        Path:    "./_support/httperf",
        Options: options,
        Parser:  true,
    }

    /* Run httperf, catching any errors with return value. */
    var output bytes.Buffer

    cmd, err := httperf.Fork(&output)

    deferred := func() {
        if err := httperf.Wait(cmd, &output); err == nil {
            /* Print selected results if successful. */
            fmt.Println("Parsed:")
            fmt.Println("-------")
            fmt.Println("QPS: ", httperf.Results.ConnectionRatePerSec)
            fmt.Println("200s:", httperf.Results.ReplyStatus2xx)
        }
    }

    if err == nil {
        defer deferred()
    }

    /* do something before calling wait */
    fmt.Println("I'm waiting...\n")

    // Output:
    //I'm waiting...
    //
    //Parsed:
    //-------
    //QPS:  4524.6
    //200s: 0
}
