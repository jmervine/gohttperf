# gohttperf

[![GoDoc](http://godoc.org/gopkg.in/jmervine/httperf.v1?status.png)](http://godoc.org/gopkg.in/jmervine/httperf.v1)    [![Build Status](https://travis-ci.org/jmervine/gohttperf.svg?branch=master)](https://travis-ci.org/jmervine/gohttperf)

#### Go Port of [HTTPerf.rb](https://github.com/jmervine/httperfrb)

> Note: This should be consider experimental, for now.

##### See [examples](_example) for detailed assertion usage.

## [Documentation](http://godoc.org/gopkg.in/jmervine/httperf.v1)

```go
import "gopkg.in/jmervine/httperf.v1"
```

### Development Notes

1. When forking, create the following symlink to support [gopkg.in](http://gopkg.in/).

```
mkdir -p $(go env GOPATH)/src/gopkg.in/jmervine
ln -s ${GIT_CLONE_PATH} $(go env GOPATH)/src/gopkg.in/jmervine/gohttperf.v1
```
