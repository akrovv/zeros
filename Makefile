.PHONY: build

build:
	go build -buildmode=plugin cmd/zeros/main.go
