package main

import (
	"context"
	"io"
	"log"
	"os"
	"strconv"

	wz "github.com/tetratelabs/wazero"
	wa "github.com/tetratelabs/wazero/api"

	gltd "github.com/takanoriyanagitani/go-load-test-dyn"

	si "github.com/takanoriyanagitani/go-load-test-dyn/tcp/http/reqsrc/seed/integer"
	sis "github.com/takanoriyanagitani/go-load-test-dyn/tcp/http/reqsrc/seed/integer/stime"

	hv1 "github.com/takanoriyanagitani/go-load-test-dyn/tcp/http/v1"
	sst "github.com/takanoriyanagitani/go-load-test-dyn/tcp/http/v1/sender/std"

	wzb "github.com/takanoriyanagitani/go-load-test-dyn/tcp/http/reqsrc/seed/external/wasm/simple/integer/integer2bytes8/rt_wazero"
)

func getenvOrAlt(key, alt string) string {
	val, ok := os.LookupEnv(key)
	switch ok {
	case true:
		return val
	default:
		return alt
	}
}

func parseIntOrAlt(s string, alt int) int {
	i, e := strconv.Atoi(s)
	switch e {
	case nil:
		return i
	default:
		return alt
	}
}

var wasmLocation string = getenvOrAlt("ENV_WASM_LOC", "./path/to/module.wasm")
var wasmFuncName string = getenvOrAlt("ENV_WASM_FNC", "seed2bytes8")

var targetUrl string = getenvOrAlt("ENV_TARGET_URL", "http://localhost/api")
var targetTyp string = getenvOrAlt("ENV_TARGET_TYP", "text/plain")

var maxWasmBytes int = parseIntOrAlt(
	getenvOrAlt("ENV_WASM_BYTES_MAX", "1048576"),
	1048576,
)

var maxLoop int = parseIntOrAlt(
	getenvOrAlt("ENV_MAX_LOOP", "16"),
	16,
)

var maxWorkers int = parseIntOrAlt(
	getenvOrAlt("ENV_MAX_WORKER", "16"),
	16,
)

var cfg wz.ModuleConfig = wz.NewModuleConfig().
	WithStartFunctions()

func compiled2instanceNew(mc wz.ModuleConfig) func(wz.Runtime) func(
	context.Context,
	wz.CompiledModule,
) (wa.Module, error) {
	return func(
		rtm wz.Runtime,
	) func(context.Context, wz.CompiledModule) (wa.Module, error) {
		return func(
			ctx context.Context,
			built wz.CompiledModule,
		) (wa.Module, error) {
			return rtm.InstantiateModule(ctx, built, mc)
		}
	}
}

var runtime2instanceDefault func(wz.Runtime) func(
	context.Context,
	wz.CompiledModule,
) (wa.Module, error) = compiled2instanceNew(cfg)

func instance2loader(instance wa.Module) (gltd.LoadSingle, error) {
	i2b8, e := wzb.IntegerToBytes8fromModule(
		instance,
		wasmFuncName,
	)
	if nil != e {
		log.Fatalf("unable to the converter: %v\n", e)
	}
	var seed2slice si.BodySource64i = func(ctx context.Context, seed int64) (
		[]byte,
		error,
	) {
		val, e := i2b8(ctx, seed)
		return val[:], e
	}

	var src si.SimpleRequestSource64i = seed2slice.ToSimpleSource(
		targetUrl,
		targetTyp,
	)

	return src.ToLoadSingle(
		sis.SeedSourceUnixtimeMicros64i,
		sst.Req2tgtSTstdHttpDefault,
		hv1.TinyResponseToSinkDiscard,
	), nil
}

func startWorker(
	ctx context.Context,
	rtm wz.Runtime,
	built wz.CompiledModule,
	reqs <-chan struct{},
	processedCnt chan<- uint64,
) {
	var built2instance func(
		context.Context,
		wz.CompiledModule,
	) (wa.Module, error) = runtime2instanceDefault(rtm)

	instance, e := built2instance(ctx, built)
	if nil != e {
		log.Printf("unable to create an instance: %v", e)
		return
	}
	defer instance.Close(ctx)

	loader, e := instance2loader(instance)
	if nil != e {
		log.Printf("unable to create a loader: %v", e)
		return
	}

	var subtot uint64 = 0
	for range reqs {
		e := loader(ctx)
		if nil != e {
			log.Fatalf("unable to call: %v\n", e)
		}
		subtot += 1
	}
	processedCnt<- subtot
}

func main() {
	var ctx context.Context = context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var rtm wz.Runtime = wz.NewRuntime(ctx)
	defer rtm.Close(ctx)

	f, e := os.Open(wasmLocation)
	if nil != e {
		log.Fatalf("unable to open wasm module(ENV_WASM_LOC): %v\n", e)
	}
	defer f.Close()

	ltd := io.LimitedReader{R: f, N: int64(maxWasmBytes)}
	wasmBytes, e := io.ReadAll(&ltd)
	if nil != e {
		log.Fatalf("unable to read wasm bytes: %v\n", e)
	}

	built, e := rtm.CompileModule(ctx, wasmBytes)
	if nil != e {
		log.Fatalf("unable to compile: %v\n", e)
	}
	defer built.Close(ctx)

	var reqs chan struct{} = make(chan struct{}, maxWorkers)

	var tot chan uint64 = make(chan uint64, maxWorkers)

	for i := 0; i < maxWorkers; i++ {
		go startWorker(
			ctx,
			rtm,
			built,
			reqs,
			tot,
		)
	}

	for i := 0; i < maxLoop; i++ {
		reqs <- struct{}{}
	}
	close(reqs)

	var gtot uint64 = 0
	for i := 0; i < maxWorkers; i++ {
		var stot uint64 = <-tot
		gtot += stot
	}

	log.Printf("total req cnt: %v\n", gtot)
}
