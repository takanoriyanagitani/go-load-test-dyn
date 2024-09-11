#!/bin/sh

export ENV_WASM_LOC=${ENV_WASM_LOC:-./path/to/module.wasm}
export ENV_WASM_FNC=${ENV_WASM_FNC:-seed2bytes8}
export ENV_WASM_BYTES_MAX=1048576

export ENV_TARGET_URL=http://localhost:9088/
export ENV_TARGET_TYP=text/plain

export ENV_MAX_LOOP=131072
export ENV_MAX_WORKER=2

time ./rt_wazero
