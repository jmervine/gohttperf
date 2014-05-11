
# gohttperf

[![GoDoc](https://godoc.org/github.com/jmervine/gohttperf?status.png)](https://godoc.org/github.com/jmervine/gohttperf)    [![Build Status](https://travis-ci.org/jmervine/gohttperf.svg?branch=master)](https://travis-ci.org/jmervine/gohttperf)

#### Go Port of [HTTPerf.rb](https://github.com/jmervine/httperfrb)

> Note: This should be consider experimental, for now.

##### See [examples](_example) for detailed assertion usage.

## [Documentation](http://godoc.org/github.com/jmervine/gohttperf)

```go
import "github.com/jmervine/gohttperf"
```

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


### Types

#### HTTPerf

```go
type HTTPerf struct {
    Options map[string]interface{}
    Path    string
    Parser  bool
    Raw     string
    Results Results
}
```




### Functions
#### Command

```go
func (h *HTTPerf) Command() string
```
Command returns (*HTTPerf).Path + (*HTTPerf).arguments()

     When:
     (*HTTPerf).Path = "httperf" // the default
     (*HTTPerf).Options = map[string]interface{}{
        "hog": true,
        "verbose": true,
        "server": "localhost"
     }
    Returns:
    "httperf --hog --verbose --server localhost"



#### Fork

```go
func (h *HTTPerf) Fork(output *bytes.Buffer) (*exec.Cmd, error)
```
Fork executes the command string returned by (*HTTPerf).Command() on the shell
in a async manner, returning the *exec.Cmd pointer to be referenced in Wait when
ready to gather results.



##### Example:
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


#### Parse

```go
func (h *HTTPerf) Parse()
```
Parse runs RawParser on a current instance of (*HTTPerf).Raw



#### Run

```go
func (h *HTTPerf) Run() error
```
Run executes the command string returned by (*HTTPerf).Command() on the shell.



##### Example:
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


#### Wait

```go
func (h *HTTPerf) Wait(cmd *exec.Cmd, output *bytes.Buffer) error
```
Wait takes the *exec.Cmd returned by Fork and waits for it to complete it's run,
on completion it sets (*HTTPerf).Raw and runs the parser if requested.



#### Results

```go
type Results struct {
    Command               string
    MaxConnectBurstLength int
    TotalConnections      int
    TotalRequests         int
    TotalReplies          int
    TotalTestDuration     float64
    ConnectionRatePerSec  float64
    ConnectionRateMsConn  float64
    ConnectionTimeMin     float64
    ConnectionTimeAvg     float64
    ConnectionTimeMax     float64
    ConnectionTimeMedian  float64
    ConnectionTimeStddev  float64
    ConnectionTimeConnect float64
    ConnectionLength      float64
    RequestRatePerSec     float64
    RequestRateMsRequest  float64
    RequestSize           float64
    ReplyRateMin          float64
    ReplyRateAvg          float64
    ReplyRateMax          float64
    ReplyRateStddev       float64
    ReplyRateSamples      int
    ReplyTimeResponse     float64
    ReplyTimeTransfer     float64
    ReplySizeHeader       float64
    ReplySizeContent      float64
    ReplySizeFooter       float64
    ReplySizeTotal        float64
    ReplyStatus1xx        int
    ReplyStatus2xx        int
    ReplyStatus3xx        int
    ReplyStatus4xx        int
    ReplyStatus5xx        int
    CPUTimeUserSec        float64
    CPUTimeSystemSec      float64
    CPUTimeUserPct        float64
    CPUTimeSystemPct      float64
    CPUTimeTotalPct       float64
    NetIoKbSec            float64
    NetIoBps              string
    ErrorsTotal           int
    ErrorsClientTimeout   int
    ErrorsSocketTimeout   int
    ErrorsConnRefused     int
    ErrorsConnReset       int
    ErrorsFdUnavail       int
    ErrorsAddrUnavail     int
    ErrorsFtabFull        int
    ErrorsOther           int
    ConnectionTimes       []float64
    ConnectionTime75Pct   float64
    ConnectionTime80Pct   float64
    ConnectionTime85Pct   float64
    ConnectionTime90Pct   float64
    ConnectionTime95Pct   float64
    ConnectionTime99Pct   float64
}
```




### Functions
#### RawParser

```go
func RawParser(raw string) Results
```
RawParser parses httperf results as printed to STDOUT on run. Push results to a
new Results struct and return it.



