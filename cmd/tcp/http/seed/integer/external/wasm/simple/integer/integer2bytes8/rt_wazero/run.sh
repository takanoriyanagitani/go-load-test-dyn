#!/bin/sh

export ENV_SEED_GRPC_ADDR=localhost:50051
export ENV_TARGET_URL=http://localhost:9088/
export ENV_POST_BODY_CONTENT_TYPE=application/octet-stream
export ENV_LOOP_MAX=1048576

time ./integer2bytes
