// file: parser.go - Parser handling for httperf raw results.

package gohttperf

// Results struct stores parsed results from (*HTTPerf).Raw
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
    CpuTimeUserSec        float64
    CpuTimeSystemSec      float64
    CpuTimeUserPct        float64
    CpuTimeSystemPct      float64
    CpuTimeTotalPct       float64
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

// Run RawParser on a current instance of (*HTTPerf)
func (this *HTTPerf) Parse() {
    this.Results = RawParser(this.Raw)
}

// Parse httperf results as printed to STDOUT on run. Push results to
// a new Results struct and return it.
func RawParser(raw string) Results {
    results := Results{}
    results.Command = toS(find(raw, "^(httperf .*)"))
    results.MaxConnectBurstLength = toI(find(raw, "Maximum connect burst length: ([0-9]*?\\.?[0-9]+)$"))
    results.TotalConnections = toI(find(raw, "^Total: connections ([0-9]*?\\.?[0-9]+) "))
    results.TotalRequests = toI(find(raw, "^Total: connections .+ requests ([0-9]*?\\.?[0-9]+) "))
    results.TotalReplies = toI(find(raw, "^Total: connections .+ replies ([0-9]*?\\.?[0-9]+) "))
    results.TotalTestDuration = toF(find(raw, "^Total: connections .+ test-duration ([0-9]*?\\.?[0-9]+) "))
    results.ConnectionRatePerSec = toF(find(raw, "^Connection rate: ([0-9]*?\\.?[0-9]+) "))
    results.ConnectionRateMsConn = toF(find(raw, "^Connection rate: .+ \\(([0-9]*?\\.?[0-9]+) ms"))
    results.ConnectionTimeMin = toF(find(raw, "^Connection time \\[ms\\]: min ([0-9]*?\\.?[0-9]+) "))
    results.ConnectionTimeAvg = toF(find(raw, "^Connection time \\[ms\\]: min .+ avg ([0-9]*?\\.?[0-9]+) "))
    results.ConnectionTimeMax = toF(find(raw, "^Connection time \\[ms\\]: min .+ max ([0-9]*?\\.?[0-9]+) "))
    results.ConnectionTimeMedian = toF(find(raw, "^Connection time \\[ms\\]: min .+ median ([0-9]*?\\.?[0-9]+) "))
    results.ConnectionTimeStddev = toF(find(raw, "^Connection time \\[ms\\]: min .+ stddev ([0-9]*?\\.?[0-9]+)$"))
    results.ConnectionTimeConnect = toF(find(raw, "^Connection time \\[ms\\]: connect ([0-9]*?\\.?[0-9]+)$"))
    results.ConnectionLength = toF(find(raw, "^Connection length \\[replies\\/conn\\]: ([0-9]*?\\.?[0-9]+)$"))
    results.RequestRatePerSec = toF(find(raw, "^Request rate: ([0-9]*?\\.?[0-9]+) req"))
    results.RequestRateMsRequest = toF(find(raw, "^Request rate: .+ \\(([0-9]*?\\.?[0-9]+) ms"))
    results.RequestSize = toF(find(raw, "^Request size \\[B\\]: ([0-9]*?\\.?[0-9]+)$"))
    results.ReplyRateMin = toF(find(raw, "^Reply rate \\[replies\\/s\\]: min ([0-9]*?\\.?[0-9]+) "))
    results.ReplyRateAvg = toF(find(raw, "^Reply rate \\[replies\\/s\\]: min .+ avg ([0-9]*?\\.?[0-9]+) "))
    results.ReplyRateMax = toF(find(raw, "^Reply rate \\[replies\\/s\\]: min .+ max ([0-9]*?\\.?[0-9]+) "))
    results.ReplyRateStddev = toF(find(raw, "^Reply rate \\[replies\\/s\\]: min .+ stddev ([0-9]*?\\.?[0-9]+) "))
    results.ReplyRateSamples = toI(find(raw, "^Reply rate \\[replies\\/s\\]: min .+ \\(([0-9]*?\\.?[0-9]+) samples"))
    results.ReplyTimeResponse = toF(find(raw, "^Reply time \\[ms\\]: response ([0-9]*?\\.?[0-9]+) "))
    results.ReplyTimeTransfer = toF(find(raw, "^Reply time \\[ms\\]: response .+ transfer ([0-9]*?\\.?[0-9]+)$"))
    results.ReplySizeHeader = toF(find(raw, "^Reply size \\[B\\]: header ([0-9]*?\\.?[0-9]+) "))
    results.ReplySizeContent = toF(find(raw, "^Reply size \\[B\\]: header .+ content ([0-9]*?\\.?[0-9]+) "))
    results.ReplySizeFooter = toF(find(raw, "^Reply size \\[B\\]: header .+ footer ([0-9]*?\\.?[0-9]+) "))
    results.ReplySizeTotal = toF(find(raw, "^Reply size \\[B\\]: header .+ \\(total ([0-9]*?\\.?[0-9]+)\\)"))
    results.ReplyStatus1xx = toI(find(raw, "^Reply status: 1xx=([0-9]*?\\.?[0-9]+) "))
    results.ReplyStatus2xx = toI(find(raw, "^Reply status: .+ 2xx=([0-9]*?\\.?[0-9]+) "))
    results.ReplyStatus3xx = toI(find(raw, "^Reply status: .+ 3xx=([0-9]*?\\.?[0-9]+) "))
    results.ReplyStatus4xx = toI(find(raw, "^Reply status: .+ 4xx=([0-9]*?\\.?[0-9]+) "))
    results.ReplyStatus5xx = toI(find(raw, "^Reply status: .+ 5xx=([0-9]*?\\.?[0-9]+)"))
    results.CpuTimeUserSec = toF(find(raw, "^CPU time \\[s\\]: user ([0-9]*?\\.?[0-9]+) "))
    results.CpuTimeUserPct = toF(find(raw, "^CPU time \\[s\\]: .+ \\(user ([0-9]*?\\.?[0-9]+)\\% "))
    results.CpuTimeSystemSec = toF(find(raw, "^CPU time \\[s\\]: .+ system ([0-9]*?\\.?[0-9]+) "))
    results.CpuTimeSystemPct = toF(find(raw, "^CPU time \\[s\\]: user .+ system .+ system ([0-9]*?\\.?[0-9]+)\\% "))
    results.CpuTimeTotalPct = toF(find(raw, "^CPU time \\[s\\]: user .+ total ([0-9]*?\\.?[0-9]+)\\%"))
    results.NetIoKbSec = toF(find(raw, "^Net I\\/O: ([0-9]*?\\.?[0-9]+) KB"))
    results.NetIoBps = toS(find(raw, "^Net I\\/O: .+ \\((.+) bps\\)"))
    results.ErrorsTotal = toI(find(raw, "^Errors: total ([0-9]*?\\.?[0-9]+) "))
    results.ErrorsClientTimeout = toI(find(raw, "^Errors: total .+ client-timo ([0-9]*?\\.?[0-9]+) "))
    results.ErrorsSocketTimeout = toI(find(raw, "^Errors: total .+ socket-timo ([0-9]*?\\.?[0-9]+) "))
    results.ErrorsConnRefused = toI(find(raw, "^Errors: total .+ connrefused ([0-9]*?\\.?[0-9]+) "))
    results.ErrorsConnReset = toI(find(raw, "^Errors: total .+ connreset ([0-9]*?\\.?[0-9]+)"))
    results.ErrorsFdUnavail = toI(find(raw, "^Errors: fd-unavail ([0-9]*?\\.?[0-9]+) "))
    results.ErrorsAddrUnavail = toI(find(raw, "^Errors: fd-unavail .+ addrunavail ([0-9]*?\\.?[0-9]+) "))
    results.ErrorsFtabFull = toI(find(raw, "^Errors: fd-unavail .+ ftab-full ([0-9]*?\\.?[0-9]+) "))
    results.ErrorsOther = toI(find(raw, "^Errors: fd-unavail .+ other ([0-9]*?\\.?[0-9]+)"))
    results.ConnectionTimes = findConnectionTimes(raw)
    results.calculatePercentiles()

    return results
}
