#!/usr/bin/env bash

# Install proto3 from source
#  brew install autoconf automake libtool
#  git clone https://github.com/google/protobuf
#  ./autogen.sh ; ./configure ; make ; make install
#
# Update protoc Go bindings via
#  go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
#
# See also
#  https://github.com/grpc/grpc-go/tree/master/examples
# Install protoc-gen-validate
# go get -u github.com/lyft/protoc-gen-validate
# cd $GOPATH/src/github.com/lyft/protoc-gen-validate
# make
# will installs PGV into $GOPATH/bin


for item in $@ ; do
		if [[ "$item" == "-vl" ]]; then
			protoc -I $GOPATH/src/github.com/lyft/protoc-gen-validate -I. --go_out=. --validate_out=lang=go:. *.proto
		else
		    protoc *.proto --go_out=plugins=grpc:.
        fi
done