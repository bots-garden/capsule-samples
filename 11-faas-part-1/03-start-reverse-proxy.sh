#!/bin/bash

# bash -c "exec -a <MyProcessName> <Command>"

capsule \
   -mode=reverse-proxy \
   -config=./config.yaml \
   -backend="memory" \
   -httpPort=8888
