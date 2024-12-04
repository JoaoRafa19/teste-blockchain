build:
	go build -o ./bin/projectx


run: build
	./bin/projectx

test:
	go test -v 	./...

all: run
	echo "Hello"

test_coverage:
	@mkdir -p coverage
	@go test -coverprofile=coverage/coverage.out ./...

cover: test_coverage
	@go tool cover -func=coverage/coverage.out > coverage/coverage-summary.txt

__coverage: _coverage	
	@rm ./coverage/coverage-summary.txt 

coverage: __coverage
	@rm ./coverage/coverage.out

_coverage: cover
	@coverage=$$(go tool cover -func=coverage/coverage.out | grep total | awk '{print substr($$3, 1, length($$3)-1)}'); \
	echo "Test coverage: $$coverage%" \

