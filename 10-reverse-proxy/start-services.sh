#!/bin/bash

# bash -c "exec -a <MyProcessName> <Command>"

MESSAGE="💊 Capsule Rocks 🚀" bash -c "exec -a hello capsule \
-wasm=./capsule-hello/hello.wasm \
-mode=http \
-httpPort=9091" &

MESSAGE="💊 Capsule Rocks 🚀" bash -c "exec -a hello capsule \
-wasm=./capsule-hello/hello.wasm \
-mode=http \
-httpPort=7071" &

# pkill -f hello

MESSAGE="💊 Capsule is Amazing 😍" bash -c "exec -a hey capsule \
-wasm=./capsule-hey/hey.wasm \
-mode=http \
-httpPort=9092" &

MESSAGE="💊 Capsule is Awesome 💚" bash -c "exec -a hola capsule \
-wasm=./capsule-hola/hola.wasm \
-mode=http \
-httpPort=9093" &

MESSAGE="💊 Capsule is Awesome 💚" bash -c "exec -a hola_orange capsule \
-wasm=./capsule-hola-orange/hola.wasm \
-mode=http \
-httpPort=6061" &

MESSAGE="💊 Capsule is Awesome 💚" bash -c "exec -a hola_yellow capsule \
-wasm=./capsule-hola-yellow/hola.wasm \
-mode=http \
-httpPort=6062" &
