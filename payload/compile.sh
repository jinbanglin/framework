#!/usr/bin/env bash

protoc -I $GOPATH/src -I. --gogofast_out=plugins=grpc:. --validate_out=lang=gogo:. *.proto
