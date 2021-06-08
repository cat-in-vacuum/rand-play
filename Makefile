GO?=go
GOTEST?=go test

.PHONY: install
install:
	$(GO) mod tidy

.PHONY: build
## build: build executable file
unit-test:
	$(GOTEST) -v -race -count=1 -cover ./generator

run:
	$(GO) run main.go