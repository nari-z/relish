  
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BINARY_NAME=gateway
BINARY_DIR=.
BINARY_AMD64_LINUX=./$(BINARY_NAME)

all: test build
build:
	$(GOBUILD) -o $(BINARY_NAME) -v
test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)

build-alpine:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_AMD64_LINUX) -v
docker-build: build-alpine
	docker build -t $(BINARY_NAME) .