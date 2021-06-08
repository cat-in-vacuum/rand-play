GO?=go
GOTEST?=go test

.PHONY: install
install:
	$(GO) mod tidy

.PHONY: unit-test
unit-test:
	$(GOTEST) -v -race -count=1 -cover ./generator

run:
	$(GO) run main.go