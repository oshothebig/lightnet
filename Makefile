BINDIR := bin

.PHONY: build
build:
	go build -o $(BINDIR)/ ./cmd/...

.PHONY: clean
clean:
	rm -rf $(BINDIR)
