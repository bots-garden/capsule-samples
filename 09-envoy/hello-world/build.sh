#!/bin/bash
tinygo build -o main.wasm -scheduler=none -target wasi ./main.go
ls -lh *.wasm
