package rtw0

import (
	"context"

	gc "google.golang.org/grpc/codes"
	gs "google.golang.org/grpc/status"

	wa "github.com/tetratelabs/wazero/api"

	i2b8 "github.com/takanoriyanagitani/go-load-test-dyn/tcp/http/reqsrc/seed/external/wasm/simple/integer/integer2bytes8"
)

type IntegerToU64 struct {
	wa.Function
}

func (u IntegerToU64) Call(ctx context.Context, input uint64) (uint64, error) {
	values, e := u.Function.Call(ctx, input)
	if nil != e {
		return 0, gs.Error(gc.Internal, "unable to call")
	}

	var lv int = len(values)
	switch lv {
	case 0:
		return 0, gs.Error(gc.Internal, "no values got")
	case 1:
		return values[0], nil
	default:
		return 0, gs.Errorf(gc.Internal, "too many values: %v", lv)
	}
}

func (u IntegerToU64) ToIntegerToBytes8() i2b8.IntegerToBytes8 {
	var i2u i2b8.IntegerToU64 = func(
		ctx context.Context,
		seed int64,
	) (uint64, error) {
		var encoded uint64 = wa.EncodeI64(seed)
		return u.Call(ctx, encoded)
	}
	return i2u.ToIntegerToBytes8()
}

func IntegerToBytes8fromModule(
	mdl wa.Module,
	name string,
) (i2b8.IntegerToBytes8, error) {
	var fnc wa.Function = mdl.ExportedFunction(name)
	if nil == fnc {
		return nil, gs.Errorf(gc.Internal, "function missing: %s", name)
	}
	return IntegerToU64{fnc}.ToIntegerToBytes8(), nil
}
