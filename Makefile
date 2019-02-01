SHELL = /bin/bash

build:
	pushd tooling && \
	go build -o generate-wallpaper && \
	popd

gen:
	./tooling/generate-wallpaper