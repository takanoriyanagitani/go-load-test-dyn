package stdclient

import (
	"bytes"
	"context"
	"io"
	"net/http"

	ph1 "github.com/takanoriyanagitani/go-load-test-dyn/load-test-dyn-proto/loadtest_dyn/http/v1"
	hv1 "github.com/takanoriyanagitani/go-load-test-dyn/tcp/http/v1"
)

type TargetToResponseStd struct {
	*http.Client
}

var TargetToResponseStdDefault TargetToResponseStd = TargetToResponseStd{
	Client: http.DefaultClient,
}

type ResponseConverter func(*http.Response) (*ph1.TinyResponse, error)

func (c TargetToResponseStd) ToRequestToTargetStPOST(
	conv ResponseConverter,
) hv1.RequestToTargetST {
	return func(
		_ctx context.Context,
		req *ph1.SimpleRequest,
	) (*ph1.TinyResponse, error) {
		var url string = req.GetUrl()
		var typ string = req.GetContentType()
		var bdy []byte = req.GetBody()
		res, e := c.Client.Post(url, typ, bytes.NewReader(bdy))
		if nil != e {
			return nil, e
		}
		return conv(res)
	}
}

var ResponseConverterDiscardBody ResponseConverter = func(
	res *http.Response,
) (*ph1.TinyResponse, error) {
	var body io.ReadCloser = res.Body
	defer body.Close()
	_, e := io.Copy(io.Discard, body)
	return &ph1.TinyResponse{
		StatusCode: int32(res.StatusCode),
	}, e
}

var Req2tgtSTstdHttpDefault hv1.RequestToTargetST = TargetToResponseStdDefault.
	ToRequestToTargetStPOST(ResponseConverterDiscardBody)
