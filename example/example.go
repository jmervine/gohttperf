package main

import (
	"fmt"
	".."
)

func main() {
    /* Define httperf arguments */
	options := map[string]interface{}{
		"server"    : "www.example.com",
		"rate"      : 5,
		"num-conns" : 10,
		"hog"       : true,
	}

    /* Create HTTPerf */
    httperf := gohttperf.HTTPerf{
        Options: options,
    }

    fmt.Println("Raw:")
    fmt.Println("-------")

    /* Run httperf, catching any errors with return value. */
    if err := httperf.Run(); err == nil {
        /* Print raw results (stdout and stderr) if successful. */
        fmt.Print(httperf.Raw)
    }

    fmt.Println("\nParsed:")
    fmt.Println("-------")

    /* Turn on Verbose mode */
    httperf.Options["verbose"] = true

    /* Turn on Parser */
    httperf.Parser = true

    /* Raw results are still available via httperf.Raw */

    /* Run httperf, catching any errors with return value. */
    if err := httperf.Run(); err == nil {
        /* Print selected results if successful. */
        fmt.Println("200s:", httperf.Results.ReplyStatus2xx)
        fmt.Println(" QPS:", httperf.Results.ConnectionRatePerSec)

        /* Print percentile information if available. */
        if len(httperf.Results.ConnectionTimes) > 0 {
            fmt.Println(" 85%:", httperf.Results.ConnectionTime85Pct)
        } else {
            fmt.Println(" 85%: n/a\n=> For verbose connection times, see: https://github.com/jmervine/httperf.")
        }
    }
}

