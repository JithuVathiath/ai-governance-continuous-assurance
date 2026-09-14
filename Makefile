GO ?= go
GOFMT ?= $(shell $(GO) env GOROOT)/bin/gofmt

.PHONY: build test test-race coverage vet security format-check demo verify clean

build:
	$(GO) build -trimpath -o bin/aegis ./cmd/aegis

test:
	$(GO) test -cover ./...

test-race:
	$(GO) test -race ./...

coverage:
	$(GO) test -race -coverprofile=coverage.out ./...
	$(GO) tool cover -func=coverage.out | awk '/^total:/ { gsub("%", "", $$3); print "Total coverage: " $$3 "%"; exit ($$3 < 85) }'

vet:
	$(GO) vet ./...

security:
	$(GO) run github.com/securego/gosec/v2/cmd/gosec@v2.28.0 ./...
	$(GO) run golang.org/x/vuln/cmd/govulncheck@v1.8.0 ./...

format-check:
	test -z "$$($(GOFMT) -l cmd internal)"

demo:
	GO_BIN=$(GO) ./scripts/build-demo.sh

verify: format-check vet coverage demo

clean:
	rm -rf bin artifacts/run coverage.out
