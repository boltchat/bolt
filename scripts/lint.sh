#!/bin/bash
LINTER_RESULT="$(gofmt -s -l .)"

if [ -n "$LINTER_RESULT" ]; then
  echo "$LINTER_RESULT"
  exit 1
fi
