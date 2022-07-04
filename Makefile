BINARY_NAME=log-viewer
TMP_DIR=/tmp

all: clean test compile build

compile:
	GOOS=freebsd GOARCH=386 go build -o bin/${BINARY_NAME}-freebsd-386 .
	GOOS=linux GOARCH=386 go build -o bin/${BINARY_NAME}-linux-386 .
	GOOS=windows GOARCH=386 go build -o bin/${BINARY_NAME}-windows-386 .
	GOOS=freebsd GOARCH=amd64 go build -o bin/${BINARY_NAME}-freebsd-amd64 .
	GOOS=darwin GOARCH=amd64 go build -o bin/${BINARY_NAME}-darwin-amd64 .
	GOOS=linux GOARCH=amd64 go build -o bin/${BINARY_NAME}-linux-amd64 .
	GOOS=windows GOARCH=amd64 go build -o bin/${BINARY_NAME}-windows-amd64 .

build:
	go build -o bin/${BINARY_NAME} .

clean:
	go clean
	rm -dfr cmd
	mkdir cmd

test:
	go test ./...

coverage:
	go test ./... -coverprofile=${TMP_DIR}/coverage.out
