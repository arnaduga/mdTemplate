#!/bin/sh

APPNAME="mdTemplate"

OS=("windows" "linux" "darwin")
ARCH=("amd64" "arm64")

for GOOS in "${OS[@]}"
do
    for GOARCH in "${ARCH[@]}"
    do
        if [ "$GOOS" = "windows" ]; then
            go build -o ./dist/${APPNAME}_${GOOS}_${GOARCH}.exe
        else
            go build -o ./dist/${APPNAME}_${GOOS}_${GOARCH}
        fi
    done
done
