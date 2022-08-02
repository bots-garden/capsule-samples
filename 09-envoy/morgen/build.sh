#!/bin/bash
tinygo build -o morgen.wasm -scheduler=none --no-debug -target wasi ./morgen.go

ls -lh *.wasm
