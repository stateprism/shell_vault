#!/usr/bin/env bash

# This script is used to release a new version of the project.
cd /src || exit

bash scripts/build.sh

rm -rf release_packages
mkdir release_packages
cp -r release/* release_packages/

cd release_packages || exit

for dir in release_*; do
    sha256sum "$dir"/* > "$dir/SHA256SUMS"
    chown -Rv ReleaseAgent:ReleaseAgent "$dir"
    tar -cvzf "$dir.tar.gz" "$dir"
done



echo Done
