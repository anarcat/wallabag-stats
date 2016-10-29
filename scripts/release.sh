#!/bin/bash

curbranch=$(git rev-parse --abbrev-ref HEAD)
latesttag=$(git describe --abbrev=0 --tags)
if [ -z "$curbranch" ]; then
  exit 1
fi
if [ -z "$latesttag" ]; then
  exit 1
fi
git checkout "$latesttag"

bin=$(basename $(pwd))
arch=$(file -b $bin | awk -F , '{ print $2  }' | xargs)
version=$(grep version version.go | awk -F \" '{ print $2  }')
releasefolder="wbgstats"

if [ -n "$bin" ]; then
  make clean build
  mkdir $releasefolder
  cp "$bin" "$releasefolder/$bin-$version"
  cp "example/config.json" "$releasefolder/"
  strip "$releasefolder/$bin-$version"
  7z a "$bin-$version-$arch.7z" "$releasefolder"
  rm -r "$releasefolder"
fi

git checkout "$curbranch"

