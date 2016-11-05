#!/bin/sh

pwd=$(pwd)
docker run -v $GOPATH:/go -w /go/${pwd#$GOPATH} -it lsegal/gorb /bin/bash
