PLUGIN_PATH := .

build_plugin:
	go build -buildmode=plugin -o ${PLUGIN_PATH}/zeros.so cmd/zeros/main.go

.PHONY: build_plugin