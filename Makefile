.PHONY: all
all: bin/shmeedia

bin/shmeedia: go.mod go.sum $(shell find . -type f -name '*.go')
	go build -o $@ ./...

venv/bin/python:
	python -m venv ./venv
	./venv/bin/pip --no-color install --upgrade pip
	./venv/bin/pip --no-color install --upgrade GitPython

.PHONY: clean
clean:
	rm -f bin/shmeedia
	rm -fr ./venv

.PHONY: test
test: bin/shmeedia venv/bin/python
	rm -f output/*
	scripts/test.py
