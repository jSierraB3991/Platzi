#! /bin/bash

export SCRIPTPATH="$( cd "$(dirname "$0")" ; pwd -P )"
file_proto=$1

if [ "$file_proto" == "" ]; then
    echo "Please especificate name of file proto"
    exit 1
elif [ ! -f "$SCRIPTPATH/$file_proto" ]; then
    echo "The file $SCRIPTPATH/$file_proto is file not found"
    exit 1
fi

first_line=$(cat $file_proto | head -1)
if [ "$first_line" != "syntax = \"proto3\";" ]; then
    echo "The file $SCRIPTPATH/$file_proto is not proto file"
    exit 1
fi


echo "RUN PROTOC TO $file_proto"
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative $file_proto
