build: clean test .PHONY
	go build -o '_pkg/httperf' -v -a -race

example: .PHONY
	cd _example; go run example.go

.PHONY:
