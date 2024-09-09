#!/bin/sh

pdir="."

files=$(
    find \
        "${pdir}" \
        -type f \
        -name '*.proto'
)

protoc \
    --proto_path="${pdir}" \
    --go_out=. \
    --go_opt=paths=source_relative \
    --go-grpc_out=. \
    --go-grpc_opt=paths=source_relative \
    ${files}
