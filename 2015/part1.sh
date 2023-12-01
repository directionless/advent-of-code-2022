#!/usr/bin/env bash

echo "( is up and ) is down"
echo "So substract"
echo ""

# This sed is macos style
cat input.txt | sed -e 's/\(.\)/\1\'$'\n/g' | sort | uniq -c

