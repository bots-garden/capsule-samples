#!/bin/bash
brew update
brew install envoy
rm core.*
envoy --version
