#!/bin/bash

versionmajor=$(grep version version.go | awk -F \" '{ print $2 }' | awk -F \. '{ print $1 }')
versionminorold=$(grep version version.go | awk -F \" '{ print $2 }' | awk -F \. '{ print $2 }')
versionminornew=$((versionminorold+1))
sed -i "s/$versionmajor\.$versionminorold/$versionmajor\.$versionminornew/" version.go

