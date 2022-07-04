BINARY_NAME=log-viewer
TMP_DIR=/tmp

all: clean test compile build

compile:
    GOOS=freebsd GOARCH=386 go build -o bin/${BINARY_NAME}-freebsd-386 main.go
    GOOS=darwin GOARCH=386 go build -o bin/${BINARY_NAME}-darwin-386 main.go
    GOOS=linux GOARCH=386 go build -o bin/${BINARY_NAME}-linux-386 main.go
    GOOS=windows GOARCH=386 go build -o bin/${BINARY_NAME}-windows-386 main.go
    GOOS=freebsd GOARCH=amd64 go build -o bin/${BINARY_NAME}-freebsd-amd64 main.go
    GOOS=darwin GOARCH=amd64 go build -o bin/${BINARY_NAME}-darwin-amd64 main.go
    GOOS=linux GOARCH=amd64 go build -o bin/${BINARY_NAME}-linux-amd64 main.go
    GOOS=windows GOARCH=amd64 go build -o bin/${BINARY_NAME}-windows-amd64 main.go

build:
	go build -o cmd/${BINARY_NAME} .

clean:
	go clean
	rm -dfr cmd
	mkdir cmd

test:
	go test ./...

coverage:
	go test ./... -coverprofile=${TMP_DIR}/coverage.out
