#!/bin/bash

bin=$(basename $(pwd))
arch=$(file -b $bin | awk -F , '{ print $2  }' | xargs)
version=$(grep version version.go | awk -F \" '{ print $2  }')

cp "$bin" "$bin-$version"
strip "$bin-$version"
7z a "$bin-$version-$arch.7z" "$bin-$version"
rm "$bin-$version"

