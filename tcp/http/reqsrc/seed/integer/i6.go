package iseed

import (
	"context"

	lt "github.com/takanoriyanagitani/go-load-test-dyn/v2"

	ph1 "github.com/takanoriyanagitani/go-load-test-dyn/v2/load-test-dyn-proto/loadtest_dyn/http/v1"
	hv1 "github.com/takanoriyanagitani/go-load-test-dyn/v2/tcp/http/v1"
)

// Gets the seed(64-bit integer).
type SeedSource64i func(context.Context) (int64, error)

// Generates the request from the seed(int64).
type SimpleRequestSource64i func(context.Context, int64) (*ph1.SimpleRequest, error)

// Generates bytes from the seed(int64).
type BodySource64i func(context.Context, int64) ([]byte, error)

func (s SimpleRequestSource64i) ToLoadSingle(
	seed SeedSource64i,
	req2tgt hv1.RequestToTargetST,
	tiny2s hv1.TinyResponseToSink,
) lt.LoadSingle {
	return func(ctx context.Context) error {
		seed64i, e := seed(ctx)
		if nil != e {
			return e
		}

		req, e := s(ctx, seed64i)
		if nil != e {
			return e
		}

		res, e := req2tgt(ctx, req)
		if nil != e {
			return e
		}

		return tiny2s(ctx, res)
	}
}

// Converts to [SimpleRequestSource64i] using the url and the content type.
func (b BodySource64i) ToSimpleSource(
	url string,
	contentType string,
) SimpleRequestSource64i {
	return func(ctx context.Context, seed64i int64) (*ph1.SimpleRequest, error) {
		body, e := b(ctx, seed64i)
		if nil != e {
			return nil, e
		}
		return &ph1.SimpleRequest{
			Url:         url,
			ContentType: contentType,
			Body:        body,
		}, nil
	}
}
