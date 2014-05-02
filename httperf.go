// file: httperf.go - Main httperf package file.

package gohttperf

import (
    "bytes"
    . "github.com/jmervine/sh"
    "os/exec"
    "strconv"
    "strings"
)

// Main HTTPerf struct, (almost) all things use this.
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
func (this *HTTPerf) arguments() string {
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
    for key, val := range this.Options {
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

// Returns (*HTTPerf).Path + (*HTTPerf).arguments()
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
func (this *HTTPerf) Command() string {
    cmd := ""

    if this.Path == "" {
        cmd += "httperf"
    } else {
        cmd += this.Path
    }

    cmd += " "
    cmd += this.arguments()
    return cmd
}

// Executes the command string returned by (*HTTPerf).Command()
// on the shell.
//
//  Returns:
//
//  Error if the command fails to execute.
func (this *HTTPerf) Run() error {
    result, err := Sh(this.Command())
    this.Raw = string(result[:])

    if this.Parser {
        // Sets this.Results to parsed Results struct.
        this.Parse()
    }

    return err
}

func (this *HTTPerf) Fork(output *bytes.Buffer) (*exec.Cmd, error) {
    return ShFork(this.Command(), output)
}

func (this *HTTPerf) Wait(cmd *exec.Cmd, output *bytes.Buffer) error {
    result, err := ShWait(cmd, output)

    if err == nil {
        this.Raw = result

        if this.Parser {
            // Sets this.Results to parsed Results struct.
            this.Parse()
        }
    }

    return err
}
