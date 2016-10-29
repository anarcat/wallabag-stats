#!/bin/bash

bin=$(basename $(pwd))
arch=$(file -b $bin | awk -F , '{ print $2  }' | xargs)
version=$(grep version version.go | awk -F \" '{ print $2  }')
curbranch=$(git rev-parse --abbrev-ref HEAD)
latesttag=$(git describe --abbrev=0 --tags)

git checkout "$latesttag"
make clean build
cp "$bin" "$bin-$version"
strip "$bin-$version"
7z a "$bin-$version-$arch.7z" "$bin-$version"
rm "$bin-$version"
git checkout "$curbranch"

