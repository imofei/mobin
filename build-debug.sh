#!/bin/bash
GIN="gin"
if [ -n "$GOPATH" ]; then
    GIN="$GOPATH/bin/gin"
fi
CGO_ENABLED=0 $GIN -p=5200 -a=5201 -b="runtime/kcs-main" -i --all --excludeDir="runtime" --buildArgs="-ldflags '-s -w'" run main.go