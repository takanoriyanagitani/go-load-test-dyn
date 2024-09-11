package i2pw0

import (
	"context"

	gc "google.golang.org/grpc/codes"
	gs "google.golang.org/grpc/status"

	wa "github.com/tetratelabs/wazero/api"

	i2p "github.com/takanoriyanagitani/go-load-test-dyn/tcp/http/reqsrc/seed/external/wasm/memory/integer/integer2page"
)

type IntegerToPage struct {
	wa.Memory
	wa.Function
	Offset uint32
	Buffer i2p.Page
}

func (p *IntegerToPage) Call(ctx context.Context, input uint64) (uint32, error) {
	values, e := p.Function.Call(ctx, input)
	if nil != e {
		return 0, e
	}
	var sz int = len(values)
	switch sz {
	case 0:
		return 0, gs.Error(gc.Internal, "no return value got")
	case 1:
		return wa.DecodeU32(values[0]), nil
	default:
		return 0, gs.Errorf(gc.Internal, "too many values got: %v", sz)
	}
}

func (p *IntegerToPage) Copy() {
	buf, _ := p.Memory.Read(p.Offset, 65536)
	copy(p.Buffer.Content[:], buf)
}

func (p *IntegerToPage) ToIntegerToPage() i2p.IntegerToPage {
	return func(ctx context.Context, seed int64) (i2p.Page, error) {
		sz, e := p.Call(ctx, wa.EncodeI64(seed))
		if nil != e {
			return p.Buffer, e
		}
		p.Copy()
		p.Buffer.Size = sz
		return p.Buffer, nil
	}
}

func ItoPageFromModule(
	ctx context.Context,
	mdl wa.Module,
	seed2pageFuncName string,
	offsetFuncName string,
) (i2p IntegerToPage, e error) {
	var mem wa.Memory = mdl.Memory()
	if nil == mem {
		return i2p, gs.Error(gc.Internal, "invalid memory")
	}

	var offset wa.Function = mdl.ExportedFunction(offsetFuncName)
	if nil == offset {
		return i2p, gs.Error(gc.Internal, "invalid offset function")
	}

	var seed2page wa.Function = mdl.ExportedFunction(seed2pageFuncName)
	if nil == seed2page {
		return i2p, gs.Error(gc.Internal, "invalid seed2page function")
	}

	values, e := offset.Call(ctx)
	if nil != e {
		return i2p, e
	}

	var sz int = len(values)
	switch sz {
	case 0:
		return i2p, gs.Error(gc.Internal, "no size got")
	case 1:
		break
	default:
		return i2p, gs.Errorf(gc.Internal, "too many values: %v", sz)
	}

	var ofst uint32 = wa.DecodeU32(values[0])

	i2p.Memory = mem
	i2p.Function = seed2page
	i2p.Offset = ofst

	return i2p, nil
}
