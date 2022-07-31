#!/bin/bash
export MESSAGE="🖐 good morning 😄"
capsule \
   -wasm=./hello.wasm \
   -mode=http \
   -httpPort=8080
