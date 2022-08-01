#!/bin/bash
COUCHBASE_CLUSTER="couchbase://localhost" \
COUCHBASE_USER="admin" \
COUCHBASE_PWD="ilovepandas" \
COUCHBASE_BUCKET="wasm-data" \
capsule \
   -wasm=./hello.wasm \
   -mode=http \
   -httpPort=7070
