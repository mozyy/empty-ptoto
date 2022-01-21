#!/bin/bash

go install .

protoc \
    --proto_path=./ \
    --go_out=. --go_opt=paths=source_relative \
    --orm_out=. --orm_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    ./config/test2.proto