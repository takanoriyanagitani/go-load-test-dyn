package hv1

import (
	"context"

	ph1 "github.com/takanoriyanagitani/go-load-test-dyn/load-test-dyn-proto/loadtest_dyn/http/v1"
)

type TinyResponseToSink func(context.Context, *ph1.TinyResponse) error

var TinyResponseToSinkDiscard TinyResponseToSink = func(
	_ context.Context,
	_ *ph1.TinyResponse,
) error {
	return nil
}
