#!/bin/bash

BCMD="go run cmd/gorbgen/main.go -build"
TCMD="ruby"

set -e
shopt -s globstar nullglob dotglob
for d in test/*/**; do
  if [ -d $d ]; then
    files=($d/*.go)
    if [ ${#files[@]} -gt 0 ]; then
      echo $BCMD $d
      $BCMD $d

      testfiles=(ext/$d/test_*.rb)
      if [ ${#testfiles[@]} -gt 0 ]; then
        echo $TCMD $testfiles
        $TCMD $testfiles
      fi
    fi
  fi
done
