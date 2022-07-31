#!/bin/bash

# before using this script, you need to serve the wasm file: python3 -m http.server 9090

export MESSAGE="🖐 good morning 😄"
capsule \
   -url=http://localhost:9090/hello.wasm \
   -wasm=./tmp/hello.wasm \
   -mode=http \
   -httpPort=8080


