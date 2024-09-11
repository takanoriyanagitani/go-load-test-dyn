package i2b8

import (
	"context"
	"encoding/binary"
)

type IntegerToBytes8 func(context.Context, int64) ([8]byte, error)

type IntegerToU64 func(context.Context, int64) (uint64, error)

func (u IntegerToU64) ToIntegerToBytes8() IntegerToBytes8 {
	return func(ctx context.Context, seed int64) ([8]byte, error) {
		var buf [8]byte
		val, e := u(ctx, seed)
		binary.LittleEndian.PutUint64(buf[:], val)
		return buf, e
	}
}
