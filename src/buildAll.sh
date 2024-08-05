#!/bin/sh

APPNAME="mdTemplate"

OS=("windows" "linux" "darwin")
ARCH=("amd64" "arm64")

for GOOS in "${OS[@]}"
do
    for GOARCH in "${ARCH[@]}"
    do
        if [ "$GOOS" = "windows" ]; then
            env GOOS=${GOOS} GOARCH=${GOARCH} go build -v -o ../dist/${APPNAME}_${GOOS}_${GOARCH}.exe
        else
            env GOOS=${GOOS} GOARCH=${GOARCH} go build -v -o ../dist/${APPNAME}_${GOOS}_${GOARCH}
        fi
    done
done