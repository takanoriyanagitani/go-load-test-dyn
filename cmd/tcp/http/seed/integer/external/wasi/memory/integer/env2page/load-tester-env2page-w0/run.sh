#!/bin/sh

export ENV_SEED_GRPC_ADDR=localhost:50051
export ENV_POST_BODY_CONTENT_TYPE=application/octet-stream

export ENV_WASM_LOC="${ENV_WASM_LOC:-./path/to/module.wasm}"
export ENV_WASM_BYTES_MAX="${ENV_WASM_BYTES_MAX:-1048576}"

export ENV_WASM_FNC="${ENV_WASM_FNC:-env2page}"
export ENV_WASM_SIZ="${ENV_WASM_SIZ:-offset64k}"

export ENV_TARGET_URL="${ENV_TARGET_URL:-http://localhost/api}"
export ENV_TARGET_TYP="${ENV_TARGET_TYP:-application/json}"

export ENV_MAX_LOOP=${ENV_MAX_LOOP:-16}
export ENV_MAX_WORKER=${ENV_MAX_WORKER:-2}

export ENV_PERF_MODE=${ENV_PERF_MODE:-tick}
export ENV_TICK_DURATION=${ENV_TICK_DURATION:-250us}

export ENV_KEYS_TO_WASI=ENV_00,ENV_01,ENV_02,ENV_FF

export ENV_00=val00
export ENV_01=val01
export ENV_02=val02
export ENV_FF=valFF

time ./load-tester-env2page-w0
