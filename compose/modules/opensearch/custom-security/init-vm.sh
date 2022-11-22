#!/bin/bash

CURRENT=`sysctl -n vm.max_map_count`;
echo "CURRENT => vm.max_map_count: $CURRENT"
DESIRED="262144";
if [ "$DESIRED" -gt "$CURRENT" ]; then
    sysctl -w vm.max_map_count=262144;
    echo "CHANGED => vm.max_map_count: $DESIRED"
fi;

CURRENT=`sysctl -n fs.file-max`;
echo "CURRENT => fs.file-max: $CURRENT"
DESIRED="65536";
if [ "$DESIRED" -gt "$CURRENT" ]; then
    sysctl -w fs.file-max=65536;
    echo "CHANGED => fs.file-max: $DESIRED"
fi;