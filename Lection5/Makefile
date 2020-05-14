# Go parameters
GOCMD=go
GOLINT=golangci-lint
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=demo

all: lint test build
build: 
	$(GOBUILD) -o $(BINARY_NAME) -v
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v	
test: 
	$(GOTEST) -v ./...
lint:
	$(GOLINT) run
cover:
	$(GOTEST) ./... -coverprofile=c.out && go tool cover -html=c.out
bench:
	$(GOTEST) ./... -bench=.	
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f c.out
run:
	$(GOBUILD) -o $(BINARY_NAME) -v 
	./$(BINARY_NAME)