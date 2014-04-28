package gohttperf

import (
    "io/ioutil"
    "os"
    "testing"
)

func TestParser(t *testing.T) {
    sample := loadSampleOutput("simple.out")
    httperf := HTTPerf{Raw: sample}
    httperf.Parse()

    results := httperf.Results
    if "httperf --server localhost" != results.Command {
        t.Error("Unexpected parse value: ", results.Command)
    }
    if 101 != results.MaxConnectBurstLength {
        t.Error("Unexpected parse value: ", results.MaxConnectBurstLength)
    }
    if 102 != results.TotalConnections {
        t.Error("Unexpected parse value: ", results.TotalConnections)
    }
    if 103 != results.TotalRequests {
        t.Error("Unexpected parse value: ", results.TotalRequests)
    }
    if 104 != results.TotalReplies {
        t.Error("Unexpected parse value: ", results.TotalReplies)
    }
    if 105.000 != results.TotalTestDuration {
        t.Error("Unexpected parse value: ", results.TotalTestDuration)
    }
    if 106.0 != results.ConnectionRatePerSec {
        t.Error("Unexpected parse value: ", results.ConnectionRatePerSec)
    }
    if 107.0 != results.ConnectionRateMsConn {
        t.Error("Unexpected parse value: ", results.ConnectionRateMsConn)
    }
    if 108.0 != results.ConnectionTimeMin {
        t.Error("Unexpected parse value: ", results.ConnectionTimeMin)
    }
    if 109.0 != results.ConnectionTimeAvg {
        t.Error("Unexpected parse value: ", results.ConnectionTimeAvg)
    }
    if 110.0 != results.ConnectionTimeMax {
        t.Error("Unexpected parse value: ", results.ConnectionTimeMax)
    }
    if 111.0 != results.ConnectionTimeMedian {
        t.Error("Unexpected parse value: ", results.ConnectionTimeMedian)
    }
    if 112.0 != results.ConnectionTimeStddev {
        t.Error("Unexpected parse value: ", results.ConnectionTimeStddev)
    }
    if 113.0 != results.ConnectionTimeConnect {
        t.Error("Unexpected parse value: ", results.ConnectionTimeConnect)
    }
    if 114.000 != results.ConnectionLength {
        t.Error("Unexpected parse value: ", results.ConnectionLength)
    }
    if 115.0 != results.RequestRatePerSec {
        t.Error("Unexpected parse value: ", results.RequestRatePerSec)
    }
    if 116.0 != results.RequestRateMsRequest {
        t.Error("Unexpected parse value: ", results.RequestRateMsRequest)
    }
    if 117.0 != results.RequestSize {
        t.Error("Unexpected parse value: ", results.RequestSize)
    }
    if 118.0 != results.ReplyRateMin {
        t.Error("Unexpected parse value: ", results.ReplyRateMin)
    }
    if 119.0 != results.ReplyRateAvg {
        t.Error("Unexpected parse value: ", results.ReplyRateAvg)
    }
    if 120.0 != results.ReplyRateMax {
        t.Error("Unexpected parse value: ", results.ReplyRateMax)
    }
    if 121.0 != results.ReplyRateStddev {
        t.Error("Unexpected parse value: ", results.ReplyRateStddev)
    }
    if 122 != results.ReplyRateSamples {
        t.Error("Unexpected parse value: ", results.ReplyRateSamples)
    }
    if 123 != results.ReplyTimeResponse {
        t.Error("Unexpected parse value: ", results.ReplyTimeResponse)
    }
    if 124 != results.ReplyTimeTransfer {
        t.Error("Unexpected parse value: ", results.ReplyTimeTransfer)
    }
    if 125.0 != results.ReplySizeHeader {
        t.Error("Unexpected parse value: ", results.ReplySizeHeader)
    }
    if 126.0 != results.ReplySizeContent {
        t.Error("Unexpected parse value: ", results.ReplySizeContent)
    }
    if 127.0 != results.ReplySizeFooter {
        t.Error("Unexpected parse value: ", results.ReplySizeFooter)
    }
    if 128.0 != results.ReplySizeTotal {
        t.Error("Unexpected parse value: ", results.ReplySizeTotal)
    }
    if 129 != results.ReplyStatus1xx {
        t.Error("Unexpected parse value: ", results.ReplyStatus1xx)
    }
    if 130 != results.ReplyStatus2xx {
        t.Error("Unexpected parse value: ", results.ReplyStatus2xx)
    }
    if 131 != results.ReplyStatus3xx {
        t.Error("Unexpected parse value: ", results.ReplyStatus3xx)
    }
    if 132 != results.ReplyStatus4xx {
        t.Error("Unexpected parse value: ", results.ReplyStatus4xx)
    }
    if 133 != results.ReplyStatus5xx {
        t.Error("Unexpected parse value: ", results.ReplyStatus5xx)
    }
    if 134.0 != results.CpuTimeUserSec {
        t.Error("Unexpected parse value: ", results.CpuTimeUserPct)
    }
    if 135.0 != results.CpuTimeSystemSec {
        t.Error("Unexpected parse value: ", results.CpuTimeSystemSec)
    }
    if 136.0 != results.CpuTimeUserPct {
        t.Error("Unexpected parse value: ", results.CpuTimeUserPct)
    }
    if 137.0 != results.CpuTimeSystemPct {
        t.Error("Unexpected parse value: ", results.CpuTimeSystemPct)
    }
    if 138.0 != results.CpuTimeTotalPct {
        t.Error("Unexpected parse value: ", results.CpuTimeTotalPct)
    }
    if 139.0 != results.NetIoKbSec {
        t.Error("Unexpected parse value: ", results.NetIoKbSec)
    }
    if "13.9*10^6" != results.NetIoBps {
        t.Error("Unexpected parse value: ", results.NetIoBps)
    }
    if 140 != results.ErrorsTotal {
        t.Error("Unexpected parse value: ", results.ErrorsTotal)
    }
    if 141 != results.ErrorsClientTimeout {
        t.Error("Unexpected parse value: ", results.ErrorsClientTimeout)
    }
    if 142 != results.ErrorsSocketTimeout {
        t.Error("Unexpected parse value: ", results.ErrorsSocketTimeout)
    }
    if 143 != results.ErrorsConnRefused {
        t.Error("Unexpected parse value: ", results.ErrorsConnRefused)
    }
    if 144 != results.ErrorsConnReset {
        t.Error("Unexpected parse value: ", results.ErrorsConnReset)
    }
    if 145 != results.ErrorsFdUnavail {
        t.Error("Unexpected parse value: ", results.ErrorsFdUnavail)
    }
    if 146 != results.ErrorsAddrUnavail {
        t.Error("Unexpected parse value: ", results.ErrorsAddrUnavail)
    }
    if 147 != results.ErrorsFtabFull {
        t.Error("Unexpected parse value: ", results.ErrorsFtabFull)
    }
    if 148 != results.ErrorsOther {
        t.Error("Unexpected parse value: ", results.ErrorsOther)
    }
}

func TestRawParser(t *testing.T) {
    sample := loadSampleOutput("simple.out")
    results := RawParser(sample)

    if "httperf --server localhost" != results.Command {
        t.Error("Unexpected parse value: ", results.Command)
    }
    if 101 != results.MaxConnectBurstLength {
        t.Error("Unexpected parse value: ", results.MaxConnectBurstLength)
    }
    if 102 != results.TotalConnections {
        t.Error("Unexpected parse value: ", results.TotalConnections)
    }
    if 103 != results.TotalRequests {
        t.Error("Unexpected parse value: ", results.TotalRequests)
    }
    if 104 != results.TotalReplies {
        t.Error("Unexpected parse value: ", results.TotalReplies)
    }
    if 105.000 != results.TotalTestDuration {
        t.Error("Unexpected parse value: ", results.TotalTestDuration)
    }
    if 106.0 != results.ConnectionRatePerSec {
        t.Error("Unexpected parse value: ", results.ConnectionRatePerSec)
    }
    if 107.0 != results.ConnectionRateMsConn {
        t.Error("Unexpected parse value: ", results.ConnectionRateMsConn)
    }
    if 108.0 != results.ConnectionTimeMin {
        t.Error("Unexpected parse value: ", results.ConnectionTimeMin)
    }
    if 109.0 != results.ConnectionTimeAvg {
        t.Error("Unexpected parse value: ", results.ConnectionTimeAvg)
    }
    if 110.0 != results.ConnectionTimeMax {
        t.Error("Unexpected parse value: ", results.ConnectionTimeMax)
    }
    if 111.0 != results.ConnectionTimeMedian {
        t.Error("Unexpected parse value: ", results.ConnectionTimeMedian)
    }
    if 112.0 != results.ConnectionTimeStddev {
        t.Error("Unexpected parse value: ", results.ConnectionTimeStddev)
    }
    if 113.0 != results.ConnectionTimeConnect {
        t.Error("Unexpected parse value: ", results.ConnectionTimeConnect)
    }
    if 114.000 != results.ConnectionLength {
        t.Error("Unexpected parse value: ", results.ConnectionLength)
    }
    if 115.0 != results.RequestRatePerSec {
        t.Error("Unexpected parse value: ", results.RequestRatePerSec)
    }
    if 116.0 != results.RequestRateMsRequest {
        t.Error("Unexpected parse value: ", results.RequestRateMsRequest)
    }
    if 117.0 != results.RequestSize {
        t.Error("Unexpected parse value: ", results.RequestSize)
    }
    if 118.0 != results.ReplyRateMin {
        t.Error("Unexpected parse value: ", results.ReplyRateMin)
    }
    if 119.0 != results.ReplyRateAvg {
        t.Error("Unexpected parse value: ", results.ReplyRateAvg)
    }
    if 120.0 != results.ReplyRateMax {
        t.Error("Unexpected parse value: ", results.ReplyRateMax)
    }
    if 121.0 != results.ReplyRateStddev {
        t.Error("Unexpected parse value: ", results.ReplyRateStddev)
    }
    if 122 != results.ReplyRateSamples {
        t.Error("Unexpected parse value: ", results.ReplyRateSamples)
    }
    if 123 != results.ReplyTimeResponse {
        t.Error("Unexpected parse value: ", results.ReplyTimeResponse)
    }
    if 124 != results.ReplyTimeTransfer {
        t.Error("Unexpected parse value: ", results.ReplyTimeTransfer)
    }
    if 125.0 != results.ReplySizeHeader {
        t.Error("Unexpected parse value: ", results.ReplySizeHeader)
    }
    if 126.0 != results.ReplySizeContent {
        t.Error("Unexpected parse value: ", results.ReplySizeContent)
    }
    if 127.0 != results.ReplySizeFooter {
        t.Error("Unexpected parse value: ", results.ReplySizeFooter)
    }
    if 128.0 != results.ReplySizeTotal {
        t.Error("Unexpected parse value: ", results.ReplySizeTotal)
    }
    if 129 != results.ReplyStatus1xx {
        t.Error("Unexpected parse value: ", results.ReplyStatus1xx)
    }
    if 130 != results.ReplyStatus2xx {
        t.Error("Unexpected parse value: ", results.ReplyStatus2xx)
    }
    if 131 != results.ReplyStatus3xx {
        t.Error("Unexpected parse value: ", results.ReplyStatus3xx)
    }
    if 132 != results.ReplyStatus4xx {
        t.Error("Unexpected parse value: ", results.ReplyStatus4xx)
    }
    if 133 != results.ReplyStatus5xx {
        t.Error("Unexpected parse value: ", results.ReplyStatus5xx)
    }
    if 134.0 != results.CpuTimeUserSec {
        t.Error("Unexpected parse value: ", results.CpuTimeUserPct)
    }
    if 135.0 != results.CpuTimeSystemSec {
        t.Error("Unexpected parse value: ", results.CpuTimeSystemSec)
    }
    if 136.0 != results.CpuTimeUserPct {
        t.Error("Unexpected parse value: ", results.CpuTimeUserPct)
    }
    if 137.0 != results.CpuTimeSystemPct {
        t.Error("Unexpected parse value: ", results.CpuTimeSystemPct)
    }
    if 138.0 != results.CpuTimeTotalPct {
        t.Error("Unexpected parse value: ", results.CpuTimeTotalPct)
    }
    if 139.0 != results.NetIoKbSec {
        t.Error("Unexpected parse value: ", results.NetIoKbSec)
    }
    if "13.9*10^6" != results.NetIoBps {
        t.Error("Unexpected parse value: ", results.NetIoBps)
    }
    if 140 != results.ErrorsTotal {
        t.Error("Unexpected parse value: ", results.ErrorsTotal)
    }
    if 141 != results.ErrorsClientTimeout {
        t.Error("Unexpected parse value: ", results.ErrorsClientTimeout)
    }
    if 142 != results.ErrorsSocketTimeout {
        t.Error("Unexpected parse value: ", results.ErrorsSocketTimeout)
    }
    if 143 != results.ErrorsConnRefused {
        t.Error("Unexpected parse value: ", results.ErrorsConnRefused)
    }
    if 144 != results.ErrorsConnReset {
        t.Error("Unexpected parse value: ", results.ErrorsConnReset)
    }
    if 145 != results.ErrorsFdUnavail {
        t.Error("Unexpected parse value: ", results.ErrorsFdUnavail)
    }
    if 146 != results.ErrorsAddrUnavail {
        t.Error("Unexpected parse value: ", results.ErrorsAddrUnavail)
    }
    if 147 != results.ErrorsFtabFull {
        t.Error("Unexpected parse value: ", results.ErrorsFtabFull)
    }
    if 148 != results.ErrorsOther {
        t.Error("Unexpected parse value: ", results.ErrorsOther)
    }

    if len(results.ConnectionTimes) != 0 {
        t.Error("Something went wrong parsing connection times.")
    }
}

func TestParserVerbose(t *testing.T) {
    sample := loadSampleOutput("verbose.out")
    httperf := HTTPerf{Raw: sample}
    httperf.Parse()

    results := httperf.Results

    if len(results.ConnectionTimes) != 100 {
        t.Error("Something went wrong parsing connection times.")
    }

    if results.ConnectionTimes[1] != 532.5 {
        t.Error("Something went wrong parsing connection times.")
    }
}
func TestRawParserVerbose(t *testing.T) {
    sample := loadSampleOutput("verbose.out")
    results := RawParser(sample)

    if len(results.ConnectionTimes) != 100 {
        t.Error("Something went wrong parsing connection times.")
    }

    if results.ConnectionTimes[1] != 532.5 {
        t.Error("Something went wrong parsing connection times.")
    }
}

// test helpers
func loadSampleOutput(f string) string {
    root, _ := os.Getwd()
    content, err := ioutil.ReadFile(root + "/test_support/" + f)
    if err != nil {
        panic(err)
    }
    return string(content)
}
