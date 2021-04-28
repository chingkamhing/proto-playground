#!/bin/bash
#
# Script to install various libraries to generate the proto files.
#

# https://github.com/protocolbuffers/protobuf/releases/
PROTOC_VER="3.15.3"
PROTOC_ZIP="protoc-${PROTOC_VER}-linux-x86_64.zip"
INSTALL_DIR="${HOME}/.local"

# either 0 argument
NUM_ARGS=0

# Function
SCRIPT_NAME=${0##*/}
Usage () {
	echo
	echo "Description:"
	echo "Script to install various libraries to generate the proto files."
	echo
	echo "Usage: $SCRIPT_NAME"
	echo "Options:"
	echo " -h                           This help message"
	echo
}

# Parse input argument(s)
while [ "${1:0:1}" == "-" ]; do
	OPT=${1:1:1}
	case "$OPT" in
	"h")
		Usage
		exit
		;;
	esac
	shift
done

if [ "$#" -ne "$NUM_ARGS" ]; then
    echo "Invalid parameter!"
	Usage
	exit 1
fi

# install protoc
curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VER}/${PROTOC_ZIP}
unzip -o $PROTOC_ZIP -d ${INSTALL_DIR} bin/protoc
unzip -o $PROTOC_ZIP -d ${INSTALL_DIR}/bin 'include/*'
rm -f $PROTOC_ZIP

# update/download necessary proto files
if [ -d ${GOPATH}/src/github.com/grpc-ecosystem ]; then
	git -C ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway pull https://github.com/grpc-ecosystem/grpc-gateway.git
else
	mkdir -p ${GOPATH}/src/github.com/grpc-ecosystem
	git -C ${GOPATH}/src/github.com/grpc-ecosystem clone https://github.com/grpc-ecosystem/grpc-gateway.git
fi
if [ -d ${GOPATH}/src/github.com/googleapis ]; then
	git -C ${GOPATH}/src/github.com/googleapis/googleapis pull https://github.com/googleapis/googleapis.git
else
	mkdir -p ${GOPATH}/src/github.com/googleapis
	git -C ${GOPATH}/src/github.com/googleapis clone https://github.com/googleapis/googleapis.git
fi

# get protocol buffer libraries
go get -u github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
go get -u google.golang.org/protobuf/cmd/protoc-gen-go
go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
go get -u github.com/favadi/protoc-go-inject-tag

# install programs for proto js generation
sudo npm install protobufjs@6.8.4 -g
sudo npm install @egret/protobuf -g
