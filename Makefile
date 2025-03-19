BUILD=./bin/xaphir

run: build
	@${BUILD}

build:
	@go build -o ${BUILD}

test:
	@go test -v -count=1 ./...

linter:
	golangci-lint run