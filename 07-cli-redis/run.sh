#!/bin/bash
REDIS_ADDR="localhost:6379" \
REDIS_PWD="" \
capsule \
   -wasm=./hello.wasm \
   -mode=cli
