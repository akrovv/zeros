PLUGIN_PATH := .

build:
	go build -buildmode=plugin -o ${PLUGIN_PATH}/zeros.so cmd/zeros/main.go

.PHONY: build