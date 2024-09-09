package segi

import (
	"context"

	rsi "github.com/takanoriyanagitani/go-load-test-dyn/load-test-dyn-proto/loadtest_dyn/raw/seed/integer/v1"
)

type ByteSource64i func(context.Context, int64) ([]byte, error)

type IndirectByteSourceGrpc struct {
	rsi.Seed64IServiceClient
}

func (i IndirectByteSourceGrpc) ToByteSource64i() ByteSource64i {
	return func(ctx context.Context, seed int64) ([]byte, error) {
		req := &rsi.BytesFrom64IRequest{Seed: seed}
		res, e := i.Seed64IServiceClient.BytesFrom64I(
			ctx,
			req,
		)
		return res.GetGenerated(), e
	}
}
