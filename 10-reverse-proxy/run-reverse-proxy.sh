#!/bin/bash
capsule \
   -mode=reverse-proxy \
   -config=./config.yaml \
   -httpPort=8888
