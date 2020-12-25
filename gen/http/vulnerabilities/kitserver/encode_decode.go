// Code generated by goa v3.2.6, DO NOT EDIT.
//
// vulnerabilities go-kit HTTP server encoders and decoders
//
// Command:
// $ goa gen citadel/design

package server

import (
	"citadel/gen/http/vulnerabilities/server"
	"context"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	goahttp "goa.design/goa/v3/http"
)

// EncodeFindResponse returns a go-kit EncodeResponseFunc suitable for encoding
// vulnerabilities find responses.
func EncodeFindResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeFindResponse(encoder)
}

// DecodeFindRequest returns a go-kit DecodeRequestFunc suitable for decoding
// vulnerabilities find requests.
func DecodeFindRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeFindRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeFindError returns a go-kit EncodeResponseFunc suitable for encoding
// errors returned by the vulnerabilities find endpoint.
func EncodeFindError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) kithttp.ErrorEncoder {
	enc := server.EncodeFindError(encoder, formatter)
	return func(ctx context.Context, err error, w http.ResponseWriter) {
		enc(ctx, w, err)
	}
}

// EncodeListResponse returns a go-kit EncodeResponseFunc suitable for encoding
// vulnerabilities list responses.
func EncodeListResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeListResponse(encoder)
}

// DecodeListRequest returns a go-kit DecodeRequestFunc suitable for decoding
// vulnerabilities list requests.
func DecodeListRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeListRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeSubmitResponse returns a go-kit EncodeResponseFunc suitable for
// encoding vulnerabilities submit responses.
func EncodeSubmitResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeSubmitResponse(encoder)
}

// DecodeSubmitRequest returns a go-kit DecodeRequestFunc suitable for decoding
// vulnerabilities submit requests.
func DecodeSubmitRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeSubmitRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeSubmitError returns a go-kit EncodeResponseFunc suitable for encoding
// errors returned by the vulnerabilities submit endpoint.
func EncodeSubmitError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) kithttp.ErrorEncoder {
	enc := server.EncodeSubmitError(encoder, formatter)
	return func(ctx context.Context, err error, w http.ResponseWriter) {
		enc(ctx, w, err)
	}
}
