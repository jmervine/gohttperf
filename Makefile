# tests without -tabs for go tip
travis: get .PHONY
	# Run Test Suite
	go test -test.v=true

build: clean test .PHONY
	go build -o 'pkg/httperf' -v -a -race

clean: .PHONY
	go clean -x

test: format .PHONY
	go test -test.v=true

format: .PHONY
	gofmt -tabs=false -tabwidth=4 -w=true -l=true *.go

get: .PHONY
	go test -i

docs: format .PHONY
	@godoc -ex=true -tabwidth=2 . | sed -e 's/func /\nfunc /g' | less
	@#                              ^ add a little spacing for readability

readme: test
	# generating readme (quietly)
	@echo '# gohttperf\n\n' > README.md
	@echo '## Use' >> README.md
	@echo '```\nimport "github.com/jmervine/gohttperf"\n```\n' >> README.md
	@echo '## Documentation\n\n```' >> README.md
	@godoc -ex=true . >> README.md
	@echo '```\n' >> README.md
	@echo '## Development\n' >> README.md
	@echo '* `make`        - run tests' >> README.md
	@echo '* `make docs`   - display godocs' >> README.md
	@echo '* `make format` - gofmt with my prefered options' >> README.md
	@echo '* `make readme` - generate README.md using godoc\n\n' >> README.md
	@# clean up whitespace
	@sed -i -e 's/\t/    /g' README.md
	@# add a little spacing for readability
	@sed -i -e 's/func /\nfunc /g' README.md

.PHONY:
