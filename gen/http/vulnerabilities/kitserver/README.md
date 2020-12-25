<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# server

```go
import "citadel/gen/http/vulnerabilities/kitserver"
```

## Index

- [func DecodeFindRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc](<#func-decodefindrequest>)
- [func DecodeListRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc](<#func-decodelistrequest>)
- [func DecodeSubmitRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc](<#func-decodesubmitrequest>)
- [func EncodeFindError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) kithttp.ErrorEncoder](<#func-encodefinderror>)
- [func EncodeFindResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc](<#func-encodefindresponse>)
- [func EncodeListResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc](<#func-encodelistresponse>)
- [func EncodeSubmitError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) kithttp.ErrorEncoder](<#func-encodesubmiterror>)
- [func EncodeSubmitResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc](<#func-encodesubmitresponse>)
- [func MountFindHandler(mux goahttp.Muxer, h http.Handler)](<#func-mountfindhandler>)
- [func MountListHandler(mux goahttp.Muxer, h http.Handler)](<#func-mountlisthandler>)
- [func MountSubmitHandler(mux goahttp.Muxer, h http.Handler)](<#func-mountsubmithandler>)


## func DecodeFindRequest

```go
func DecodeFindRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc
```

DecodeFindRequest returns a go\-kit DecodeRequestFunc suitable for decoding vulnerabilities find requests\.

## func DecodeListRequest

```go
func DecodeListRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc
```

DecodeListRequest returns a go\-kit DecodeRequestFunc suitable for decoding vulnerabilities list requests\.

## func DecodeSubmitRequest

```go
func DecodeSubmitRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc
```

DecodeSubmitRequest returns a go\-kit DecodeRequestFunc suitable for decoding vulnerabilities submit requests\.

## func EncodeFindError

```go
func EncodeFindError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) kithttp.ErrorEncoder
```

EncodeFindError returns a go\-kit EncodeResponseFunc suitable for encoding errors returned by the vulnerabilities find endpoint\.

## func EncodeFindResponse

```go
func EncodeFindResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc
```

EncodeFindResponse returns a go\-kit EncodeResponseFunc suitable for encoding vulnerabilities find responses\.

## func EncodeListResponse

```go
func EncodeListResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc
```

EncodeListResponse returns a go\-kit EncodeResponseFunc suitable for encoding vulnerabilities list responses\.

## func EncodeSubmitError

```go
func EncodeSubmitError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) kithttp.ErrorEncoder
```

EncodeSubmitError returns a go\-kit EncodeResponseFunc suitable for encoding errors returned by the vulnerabilities submit endpoint\.

## func EncodeSubmitResponse

```go
func EncodeSubmitResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc
```

EncodeSubmitResponse returns a go\-kit EncodeResponseFunc suitable for encoding vulnerabilities submit responses\.

## func MountFindHandler

```go
func MountFindHandler(mux goahttp.Muxer, h http.Handler)
```

MountFindHandler configures the mux to serve the "vulnerabilities" service "find" endpoint\.

## func MountListHandler

```go
func MountListHandler(mux goahttp.Muxer, h http.Handler)
```

MountListHandler configures the mux to serve the "vulnerabilities" service "list" endpoint\.

## func MountSubmitHandler

```go
func MountSubmitHandler(mux goahttp.Muxer, h http.Handler)
```

MountSubmitHandler configures the mux to serve the "vulnerabilities" service "submit" endpoint\.



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)