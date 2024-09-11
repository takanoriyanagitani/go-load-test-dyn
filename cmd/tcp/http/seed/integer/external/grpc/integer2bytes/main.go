package main

import (
	"context"
	"os"
	"strconv"

	"google.golang.org/grpc"
	gc "google.golang.org/grpc/credentials"
	gci "google.golang.org/grpc/credentials/insecure"

	gif "github.com/takanoriyanagitani/go-load-test-dyn/load-test-dyn-proto/loadtest_dyn/raw/seed/integer/v1"

	egi "github.com/takanoriyanagitani/go-load-test-dyn/tcp/http/reqsrc/seed/external/grpc/integer"

	rsi "github.com/takanoriyanagitani/go-load-test-dyn/tcp/http/reqsrc/seed/integer"
	sis "github.com/takanoriyanagitani/go-load-test-dyn/tcp/http/reqsrc/seed/integer/stime"

	hv1 "github.com/takanoriyanagitani/go-load-test-dyn/tcp/http/v1"
	sst "github.com/takanoriyanagitani/go-load-test-dyn/tcp/http/v1/sender/std"

	gltd "github.com/takanoriyanagitani/go-load-test-dyn"
)

func must[T any](t T, e error) T {
	if nil == e {
		return t
	}
	panic(e)
}

func getenvOrAlt(key, alt string) string {
	val, ok := os.LookupEnv(key)
	switch ok {
	case true:
		return val
	default:
		return alt
	}
}

var seedGrpcAddr string = getenvOrAlt(
	"ENV_SEED_GRPC_ADDR",
	"localhost:50051",
)

var targetUrl string = getenvOrAlt(
	"ENV_TARGET_URL",
	"http://localhost:8080/",
)

var postBodyContentType string = getenvOrAlt(
	"ENV_POST_BODY_CONTENT_TYPE",
	"application/octet-stream",
)

var loopMax string = getenvOrAlt(
	"ENV_LOOP_MAX",
	"10",
)

var numOfLoops int = must(strconv.Atoi(loopMax))

var tc gc.TransportCredentials = gci.NewCredentials()

var conn *grpc.ClientConn = must(grpc.NewClient(
	seedGrpcAddr,
	grpc.WithTransportCredentials(tc),
))
var cci grpc.ClientConnInterface = conn

var sclient gif.Seed64IServiceClient = gif.NewSeed64IServiceClient(
	cci,
)

var bsource egi.ByteSource64i = egi.IndirectByteSourceGrpc{
	Seed64IServiceClient: sclient,
}.
	ToByteSource64i()
var bodysrc rsi.BodySource64i = rsi.BodySource64i(bsource)
var reqsrc rsi.SimpleRequestSource64i = bodysrc.
	ToSimpleSource(
		targetUrl,
		postBodyContentType,
	)

var seedSource rsi.SeedSource64i = sis.SeedSourceUnixtimeMicros64i

var req2tgt hv1.RequestToTargetST = sst.Req2tgtSTstdHttpDefault
var res2snk hv1.TinyResponseToSink = hv1.TinyResponseToSinkDiscard

var call1 gltd.LoadSingle = reqsrc.
	ToLoadSingle(
		seedSource,
		req2tgt,
		res2snk,
	)

func main() {
	var ctx context.Context = context.Background()

	for range numOfLoops {
		e := call1(ctx)
		if nil != e {
			panic(e)
		}
	}
}
