#!/bin/bash
#tinygo build -o hello.wasm --no-debug -target wasi ./hello.go
tinygo build -o hello.wasm -scheduler=none --no-debug -target wasi ./hello.go
#tinygo build -o hello.wasm -scheduler=none -target wasi ./hello.go


ls -lh *.wasm
