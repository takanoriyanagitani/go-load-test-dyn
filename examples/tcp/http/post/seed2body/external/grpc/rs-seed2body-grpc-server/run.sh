#!/bin/sh

export ENV_LISTEN_ADDR=127.0.0.1:50051
./target/release/rs-seed2body-grpc-server
