// Code generated by goa v3.2.6, DO NOT EDIT.
//
// incidents client HTTP transport
//
// Command:
// $ goa gen github.com/brittonhayes/citadel/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the incidents service endpoint HTTP clients.
type Client struct {
	// Find Doer is the HTTP client used to make requests to the find endpoint.
	FindDoer goahttp.Doer

	// ListAll Doer is the HTTP client used to make requests to the list all
	// endpoint.
	ListAllDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the incidents service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		FindDoer:            doer,
		ListAllDoer:         doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Find returns an endpoint that makes HTTP requests to the incidents service
// find server.
func (c *Client) Find() goa.Endpoint {
	var (
		decodeResponse = DecodeFindResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildFindRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.FindDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("incidents", "find", err)
		}
		return decodeResponse(resp)
	}
}

// ListAll returns an endpoint that makes HTTP requests to the incidents
// service list all server.
func (c *Client) ListAll() goa.Endpoint {
	var (
		encodeRequest  = EncodeListAllRequest(c.encoder)
		decodeResponse = DecodeListAllResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildListAllRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ListAllDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("incidents", "list all", err)
		}
		return decodeResponse(resp)
	}
}
