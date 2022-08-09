#!/bin/bash

# bash -c "exec -a <MyProcessName> <Command>"

capsule \
   -mode=registry \
   -files="./functions" \
   -httpPort=4999
