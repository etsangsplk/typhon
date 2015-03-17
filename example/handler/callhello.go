package handler

import (
	"fmt"

	"github.com/b2aio/typhon/example/proto/callhello"
	"github.com/b2aio/typhon/example/proto/hello"
	"github.com/b2aio/typhon/server"
)

var CallHello = &server.Endpoint{
	Name:     "callhello",
	Handler:  callhelloHandler,
	Request:  &callhello.Request{},
	Response: &callhello.Response{},
}

func callhelloHandler(req server.Request) (server.Response, error) {

	reqProto := req.Body().(*callhello.Request)
	resp := &hello.Response{}

	req.ScopedRequest(
		"example",
		"hello",
		&hello.Request{
			Name: reqProto.Value,
		},
		resp,
	)

	return server.NewProtoResponse(&callhello.Response{
		Value: fmt.Sprintf("example.hello says '%s'", resp.Greeting),
	}), nil
}
