#!/bin/sh

target=localhost:50051

grpcurl \
	-plaintext \
	-import-path ./load-test-dyn-proto \
	-proto loadtest_dyn/raw/seed/integer/v1/seed2bytes.proto \
	-d '{
		"seed": 124
	}' \
	"${target}" \
	loadtest_dyn.raw.seed.integer.v1.Seed64iService/BytesFrom64i
