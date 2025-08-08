#!/usr/bin/env bash

INPUT="$1"
OUTPUT="$2"

cp "$INPUT" "$OUTPUT" || exit 1

mogrify -format jxl -strip "$OUTPUT" || exit 2

QUALITY=$(identify -format %Q $OUTPUT)

if [ $QUALITY -gt 90 ]; then
  mogrify -quality 90 "$OUTPUT" || exit 4
fi
