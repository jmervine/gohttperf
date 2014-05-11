# tests without -tabs for go tip
travis: get .PHONY
	# Run Test Suite
	go test -test.v=true

build: clean test .PHONY
	go build -o '_pkg/httperf' -v -a -race

clean: .PHONY
	go clean -x

test: format .PHONY
	go test -test.v=true

quiet/test: .PHONY
	go test

example: .PHONY
	cd _example; go run example.go

format: .PHONY
	gofmt -tabs=false -tabwidth=4 -w=true -l=true *.go

get: .PHONY
	go get github.com/jmervine/sh

docs: format .PHONY
	@godoc -ex=true -tabwidth=2 . | sed -e 's/func /\nfunc /g' | less
	@#                              ^ add a little spacing for readability

readme: format quiet/test .PHONY
	# generating readme
	godoc -ex -v -templates "$(PWD)/_support" . > README.md

.PHONY:
