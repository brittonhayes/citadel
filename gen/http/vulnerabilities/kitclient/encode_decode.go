// Code generated by goa v3.2.6, DO NOT EDIT.
//
// vulnerabilities go-kit HTTP client encoders and decoders
//
// Command:
// $ goa gen citadel/design

package client

import (
	"citadel/gen/http/vulnerabilities/client"
	"context"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	goahttp "goa.design/goa/v3/http"
)

// DecodeFindResponse returns a go-kit DecodeResponseFunc suitable for decoding
// vulnerabilities find responses.
func DecodeFindResponse(decoder func(*http.Response) goahttp.Decoder) kithttp.DecodeResponseFunc {
	dec := client.DecodeFindResponse(decoder, false)
	return func(ctx context.Context, resp *http.Response) (interface{}, error) {
		return dec(resp)
	}
}

// EncodeListRequest returns a go-kit EncodeRequestFunc suitable for encoding
// vulnerabilities list requests.
func EncodeListRequest(encoder func(*http.Request) goahttp.Encoder) kithttp.EncodeRequestFunc {
	enc := client.EncodeListRequest(encoder)
	return func(_ context.Context, r *http.Request, v interface{}) error {
		return enc(r, v)
	}
}

// DecodeListResponse returns a go-kit DecodeResponseFunc suitable for decoding
// vulnerabilities list responses.
func DecodeListResponse(decoder func(*http.Response) goahttp.Decoder) kithttp.DecodeResponseFunc {
	dec := client.DecodeListResponse(decoder, false)
	return func(ctx context.Context, resp *http.Response) (interface{}, error) {
		return dec(resp)
	}
}

// EncodeSubmitRequest returns a go-kit EncodeRequestFunc suitable for encoding
// vulnerabilities submit requests.
func EncodeSubmitRequest(encoder func(*http.Request) goahttp.Encoder) kithttp.EncodeRequestFunc {
	enc := client.EncodeSubmitRequest(encoder)
	return func(_ context.Context, r *http.Request, v interface{}) error {
		return enc(r, v)
	}
}

// DecodeSubmitResponse returns a go-kit DecodeResponseFunc suitable for
// decoding vulnerabilities submit responses.
func DecodeSubmitResponse(decoder func(*http.Response) goahttp.Decoder) kithttp.DecodeResponseFunc {
	dec := client.DecodeSubmitResponse(decoder, false)
	return func(ctx context.Context, resp *http.Response) (interface{}, error) {
		return dec(resp)
	}
}