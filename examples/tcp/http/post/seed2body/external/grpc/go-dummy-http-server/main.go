package main

import (
	"log"

	"net/http"
	"time"
)

var listenAddr string = ":9088"

var hserver = &http.Server{
	Addr:           listenAddr,
	ReadTimeout:    10 * time.Second,
	WriteTimeout:   10 * time.Second,
	MaxHeaderBytes: 1 << 20,
}

type SimpleMetrics interface {
	IncRequestCount()
	GetRequestCount() uint64
}

type IncRequestCountReq struct{ Reply chan struct{} }
type GetRequestCountReq struct{ Reply chan uint64 }

type SimpleMetRequests struct {
	IncRequests chan IncRequestCountReq
	GetRequests chan GetRequestCountReq
}

type SimpleMetricsState struct {
	RequestCount uint64
}

func startSimpleMetrics(reqs SimpleMetRequests) {
	var state SimpleMetricsState
	for {
		select {
		case ireq := <-reqs.IncRequests:
			(func() {
				defer close(ireq.Reply)
				state.RequestCount += 1
				if 0 == (state.RequestCount % 16384) {
					log.Printf("requests: %v\n", state.RequestCount)
				}
				ireq.Reply <- struct{}{}
			})()
		case greq := <-reqs.GetRequests:
			(func() {
				defer close(greq.Reply)
				var cnt uint64 = state.RequestCount
				greq.Reply <- cnt
			})()
		}
	}
}

type SimpleMetImpl struct {
	Sender SimpleMetRequests
}

func (s SimpleMetImpl) IncRequestCount() {
	Reply := make(chan struct{})
	ireq := IncRequestCountReq{Reply}
	s.Sender.IncRequests <- ireq
	<-ireq.Reply
}

func (s SimpleMetImpl) GetRequestCount() uint64 {
	Reply := make(chan uint64)
	greq := GetRequestCountReq{Reply}
	s.Sender.GetRequests <- greq
	var cnt uint64 = <-greq.Reply
	return cnt
}

func (s SimpleMetImpl) AsMetrics() SimpleMetrics { return s }

func newSimpleMetrics() SimpleMetrics {
	IntRequests := make(chan IncRequestCountReq)
	GetRequests := make(chan GetRequestCountReq)
	reqs := SimpleMetRequests{
		IntRequests,
		GetRequests,
	}
	go startSimpleMetrics(reqs)
	return SimpleMetImpl{Sender: reqs}.AsMetrics()
}

func main() {
	var smet SimpleMetrics = newSimpleMetrics()

	discardHandler := func(_ http.ResponseWriter, _ *http.Request) {
		smet.IncRequestCount()
	}

	http.HandleFunc("/", discardHandler)

	e := hserver.ListenAndServe()
	if nil != e {
		panic(e)
	}
}
