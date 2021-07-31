BINDIR := bin

# Build static-linked binary
export CGO_ENABLED=0

.PHONY: build
build:
	go build -o $(BINDIR)/ ./cmd/...

.PHONY: test
test:
	go test -v ./...

.PHONY: format
format:
	go fmt ./...

.PHONY: clean
clean:
	rm -rf $(BINDIR)
