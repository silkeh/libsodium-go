#!/bin/bash
env CGO_CPPFLAGS=-I/usr/include/ CGO_LDFLAGS='/usr/lib/libsodium.so' go build ./...
