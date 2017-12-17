all: build

build: deps cmd/nbx/main.go
	@\
	CGO_ENABLED=0 \
	GOOS=linux \
	go build \
		-o bin/nbx \
		-a -tags netgo \
		-ldflags '-w' \
		-v \
		cmd/nbx/main.go

dev: cmd/nbx/main.go
	@go build -o bin/nbx -v cmd/nbx/*.go
	@chmod +x bin/nbx

deps:
	@dep ensure || \
		go get -u github.com/golang/dep/cmd/dep && \
		dep ensure

fmt:
	@go fmt $(go list ./... | grep -v /vendor/)
		
test:
	@go test -cover $(go list ./... | grep -v /vendor/)

guard-%:
	@ if [ "${${*}}" = "" ]; then \
	echo "Environment variable $* required for command is not set"; \
	exit 1; \
    fi
