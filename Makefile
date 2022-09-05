bin/shmeedia: go.mod go.sum $(shell find . -type f -name '*.go')
	go build -o $@ ./...

.PHONY: clean
clean:
	rm -f bin/shmeedia

.PHONY: test
test: bin/shmeedia
	rm -f output/*
	bin/shmeedia
