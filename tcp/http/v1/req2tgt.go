package hv1

import (
	"context"

	ph1 "github.com/takanoriyanagitani/go-load-test-dyn/load-test-dyn-proto/loadtest_dyn/http/v1"
)

// Tries to get the response after sending the request.
type RequestToTargetST func(
	context.Context,
	*ph1.SimpleRequest,
) (*ph1.TinyResponse, error)
