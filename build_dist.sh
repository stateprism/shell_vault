#!/usr/bin/env bash


UPX_FLAGS=("--lzma")
GOOS_LIST=("windows" "linux" "darwin")

make clean
rm -rf release
mkdir release

for os in "${GOOS_LIST[@]}"; do
    echo "Building for $os"
    GOOS=$os make build
    mkdir "release_$os"
    mv bin/* "release_$os"
    upx "${UPX_FLAGS[@]}" release_$os/*
done

mv release_* release
