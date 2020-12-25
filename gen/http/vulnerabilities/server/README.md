<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# server

```go
import "citadel/gen/http/vulnerabilities/server"
```

## Index

- [func DecodeFindRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error)](<#func-decodefindrequest>)
- [func DecodeListRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error)](<#func-decodelistrequest>)
- [func DecodeSubmitRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error)](<#func-decodesubmitrequest>)
- [func EncodeFindError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error](<#func-encodefinderror>)
- [func EncodeFindResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error](<#func-encodefindresponse>)
- [func EncodeListResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error](<#func-encodelistresponse>)
- [func EncodeSubmitError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error](<#func-encodesubmiterror>)
- [func EncodeSubmitResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error](<#func-encodesubmitresponse>)
- [func FindVulnerabilitiesPath(id uint64) string](<#func-findvulnerabilitiespath>)
- [func ListVulnerabilitiesPath() string](<#func-listvulnerabilitiespath>)
- [func Mount(mux goahttp.Muxer, h *Server)](<#func-mount>)
- [func MountFindHandler(mux goahttp.Muxer, h http.Handler)](<#func-mountfindhandler>)
- [func MountListHandler(mux goahttp.Muxer, h http.Handler)](<#func-mountlisthandler>)
- [func MountSubmitHandler(mux goahttp.Muxer, h http.Handler)](<#func-mountsubmithandler>)
- [func NewFindHandler(
    endpoint endpoint.Endpoint,
    mux goahttp.Muxer,
    decoder func(*http.Request) goahttp.Decoder,
    encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
    errhandler func(context.Context, http.ResponseWriter, error),
    formatter func(err error) goahttp.Statuser,
) http.Handler](<#func-newfindhandler>)
- [func NewFindPayload(id uint64) *vulnerabilities.FindPayload](<#func-newfindpayload>)
- [func NewListHandler(
    endpoint endpoint.Endpoint,
    mux goahttp.Muxer,
    decoder func(*http.Request) goahttp.Decoder,
    encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
    errhandler func(context.Context, http.ResponseWriter, error),
    formatter func(err error) goahttp.Statuser,
) http.Handler](<#func-newlisthandler>)
- [func NewListLimitPayload(limit *int32) *vulnerabilities.LimitPayload](<#func-newlistlimitpayload>)
- [func NewSubmitHandler(
    endpoint endpoint.Endpoint,
    mux goahttp.Muxer,
    decoder func(*http.Request) goahttp.Decoder,
    encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
    errhandler func(context.Context, http.ResponseWriter, error),
    formatter func(err error) goahttp.Statuser,
) http.Handler](<#func-newsubmithandler>)
- [func NewSubmitPayload(body *SubmitRequestBody) *vulnerabilities.SubmitPayload](<#func-newsubmitpayload>)
- [func SubmitVulnerabilitiesPath() string](<#func-submitvulnerabilitiespath>)
- [type ErrorNamer](<#type-errornamer>)
- [type FindNoMatchResponseBody](<#type-findnomatchresponsebody>)
  - [func NewFindNoMatchResponseBody(res vulnerabilities.NoMatch) FindNoMatchResponseBody](<#func-newfindnomatchresponsebody>)
- [type FindResponseBody](<#type-findresponsebody>)
  - [func NewFindResponseBody(res *vulnerabilities.Vulnerability) *FindResponseBody](<#func-newfindresponsebody>)
- [type ListResponseBody](<#type-listresponsebody>)
  - [func NewListResponseBody(res []*vulnerabilities.Vulnerability) ListResponseBody](<#func-newlistresponsebody>)
- [type MountPoint](<#type-mountpoint>)
- [type Server](<#type-server>)
  - [func New(
    e *vulnerabilities.Endpoints,
    mux goahttp.Muxer,
    decoder func(*http.Request) goahttp.Decoder,
    encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
    errhandler func(context.Context, http.ResponseWriter, error),
    formatter func(err error) goahttp.Statuser,
) *Server](<#func-new>)
  - [func (s *Server) Service() string](<#func-server-service>)
  - [func (s *Server) Use(m func(http.Handler) http.Handler)](<#func-server-use>)
- [type SubmitNoMatchResponseBody](<#type-submitnomatchresponsebody>)
  - [func NewSubmitNoMatchResponseBody(res vulnerabilities.NoMatch) SubmitNoMatchResponseBody](<#func-newsubmitnomatchresponsebody>)
- [type SubmitRequestBody](<#type-submitrequestbody>)
- [type VulnerabilityResponse](<#type-vulnerabilityresponse>)


## func DecodeFindRequest

```go
func DecodeFindRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error)
```

DecodeFindRequest returns a decoder for requests sent to the vulnerabilities find endpoint\.

## func DecodeListRequest

```go
func DecodeListRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error)
```

DecodeListRequest returns a decoder for requests sent to the vulnerabilities list endpoint\.

## func DecodeSubmitRequest

```go
func DecodeSubmitRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error)
```

DecodeSubmitRequest returns a decoder for requests sent to the vulnerabilities submit endpoint\.

## func EncodeFindError

```go
func EncodeFindError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error
```

EncodeFindError returns an encoder for errors returned by the find vulnerabilities endpoint\.

## func EncodeFindResponse

```go
func EncodeFindResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error
```

EncodeFindResponse returns an encoder for responses returned by the vulnerabilities find endpoint\.

## func EncodeListResponse

```go
func EncodeListResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error
```

EncodeListResponse returns an encoder for responses returned by the vulnerabilities list endpoint\.

## func EncodeSubmitError

```go
func EncodeSubmitError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error
```

EncodeSubmitError returns an encoder for errors returned by the submit vulnerabilities endpoint\.

## func EncodeSubmitResponse

```go
func EncodeSubmitResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error
```

EncodeSubmitResponse returns an encoder for responses returned by the vulnerabilities submit endpoint\.

## func FindVulnerabilitiesPath

```go
func FindVulnerabilitiesPath(id uint64) string
```

FindVulnerabilitiesPath returns the URL path to the vulnerabilities service find HTTP endpoint\.

## func ListVulnerabilitiesPath

```go
func ListVulnerabilitiesPath() string
```

ListVulnerabilitiesPath returns the URL path to the vulnerabilities service list HTTP endpoint\.

## func Mount

```go
func Mount(mux goahttp.Muxer, h *Server)
```

Mount configures the mux to serve the vulnerabilities endpoints\.

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

## func NewFindHandler

```go
func NewFindHandler(
    endpoint endpoint.Endpoint,
    mux goahttp.Muxer,
    decoder func(*http.Request) goahttp.Decoder,
    encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
    errhandler func(context.Context, http.ResponseWriter, error),
    formatter func(err error) goahttp.Statuser,
) http.Handler
```

NewFindHandler creates a HTTP handler which loads the HTTP request and calls the "vulnerabilities" service "find" endpoint\.

## func NewFindPayload

```go
func NewFindPayload(id uint64) *vulnerabilities.FindPayload
```

NewFindPayload builds a vulnerabilities service find endpoint payload\.

## func NewListHandler

```go
func NewListHandler(
    endpoint endpoint.Endpoint,
    mux goahttp.Muxer,
    decoder func(*http.Request) goahttp.Decoder,
    encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
    errhandler func(context.Context, http.ResponseWriter, error),
    formatter func(err error) goahttp.Statuser,
) http.Handler
```

NewListHandler creates a HTTP handler which loads the HTTP request and calls the "vulnerabilities" service "list" endpoint\.

## func NewListLimitPayload

```go
func NewListLimitPayload(limit *int32) *vulnerabilities.LimitPayload
```

NewListLimitPayload builds a vulnerabilities service list endpoint payload\.

## func NewSubmitHandler

```go
func NewSubmitHandler(
    endpoint endpoint.Endpoint,
    mux goahttp.Muxer,
    decoder func(*http.Request) goahttp.Decoder,
    encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
    errhandler func(context.Context, http.ResponseWriter, error),
    formatter func(err error) goahttp.Statuser,
) http.Handler
```

NewSubmitHandler creates a HTTP handler which loads the HTTP request and calls the "vulnerabilities" service "submit" endpoint\.

## func NewSubmitPayload

```go
func NewSubmitPayload(body *SubmitRequestBody) *vulnerabilities.SubmitPayload
```

NewSubmitPayload builds a vulnerabilities service submit endpoint payload\.

## func SubmitVulnerabilitiesPath

```go
func SubmitVulnerabilitiesPath() string
```

SubmitVulnerabilitiesPath returns the URL path to the vulnerabilities service submit HTTP endpoint\.

## type ErrorNamer

ErrorNamer is an interface implemented by generated error structs that exposes the name of the error as defined in the design\.

```go
type ErrorNamer interface {
    ErrorName() string
}
```

## type FindNoMatchResponseBody

FindNoMatchResponseBody is the type of the "vulnerabilities" service "find" endpoint HTTP response body for the "no\_match" error\.

```go
type FindNoMatchResponseBody string
```

### func NewFindNoMatchResponseBody

```go
func NewFindNoMatchResponseBody(res vulnerabilities.NoMatch) FindNoMatchResponseBody
```

NewFindNoMatchResponseBody builds the HTTP response body from the result of the "find" endpoint of the "vulnerabilities" service\.

## type FindResponseBody

FindResponseBody is the type of the "vulnerabilities" service "find" endpoint HTTP response body\.

```go
type FindResponseBody struct {
    // Unique ID of the vulnerability
    ID  uint64 `form:"id" json:"id" xml:"id"`
    // Title of the vulnerability
    Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
    // Description of the vulnerability
    Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
    // If the vulnerability is exploitable
    Exploitable *bool `form:"exploitable,omitempty" json:"exploitable,omitempty" xml:"exploitable,omitempty"`
    // Severity score of the vulnerability
    CvssScore *float32 `form:"cvss_score,omitempty" json:"cvss_score,omitempty" xml:"cvss_score,omitempty"`
    // If the vulnerability is patchable
    IsPatchable *bool `form:"is_patchable,omitempty" json:"is_patchable,omitempty" xml:"is_patchable,omitempty"`
    // If the vulnerability is upgradeable
    IsUpgradeable *bool `form:"is_upgradeable,omitempty" json:"is_upgradeable,omitempty" xml:"is_upgradeable,omitempty"`
}
```

### func NewFindResponseBody

```go
func NewFindResponseBody(res *vulnerabilities.Vulnerability) *FindResponseBody
```

NewFindResponseBody builds the HTTP response body from the result of the "find" endpoint of the "vulnerabilities" service\.

## type ListResponseBody

ListResponseBody is the type of the "vulnerabilities" service "list" endpoint HTTP response body\.

```go
type ListResponseBody []*VulnerabilityResponse
```

### func NewListResponseBody

```go
func NewListResponseBody(res []*vulnerabilities.Vulnerability) ListResponseBody
```

NewListResponseBody builds the HTTP response body from the result of the "list" endpoint of the "vulnerabilities" service\.

## type MountPoint

MountPoint holds information about the mounted endpoints\.

```go
type MountPoint struct {
    // Method is the name of the service method served by the mounted HTTP handler.
    Method string
    // Verb is the HTTP method used to match requests to the mounted handler.
    Verb string
    // Pattern is the HTTP request path pattern used to match requests to the
    // mounted handler.
    Pattern string
}
```

## type Server

Server lists the vulnerabilities service endpoint HTTP handlers\.

```go
type Server struct {
    Mounts []*MountPoint
    Find   http.Handler
    List   http.Handler
    Submit http.Handler
}
```

### func New

```go
func New(
    e *vulnerabilities.Endpoints,
    mux goahttp.Muxer,
    decoder func(*http.Request) goahttp.Decoder,
    encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
    errhandler func(context.Context, http.ResponseWriter, error),
    formatter func(err error) goahttp.Statuser,
) *Server
```

New instantiates HTTP handlers for all the vulnerabilities service endpoints using the provided encoder and decoder\. The handlers are mounted on the given mux using the HTTP verb and path defined in the design\. errhandler is called whenever a response fails to be encoded\. formatter is used to format errors returned by the service methods prior to encoding\. Both errhandler and formatter are optional and can be nil\.

### func \(\*Server\) Service

```go
func (s *Server) Service() string
```

Service returns the name of the service served\.

### func \(\*Server\) Use

```go
func (s *Server) Use(m func(http.Handler) http.Handler)
```

Use wraps the server handlers with the given middleware\.

## type SubmitNoMatchResponseBody

SubmitNoMatchResponseBody is the type of the "vulnerabilities" service "submit" endpoint HTTP response body for the "no\_match" error\.

```go
type SubmitNoMatchResponseBody string
```

### func NewSubmitNoMatchResponseBody

```go
func NewSubmitNoMatchResponseBody(res vulnerabilities.NoMatch) SubmitNoMatchResponseBody
```

NewSubmitNoMatchResponseBody builds the HTTP response body from the result of the "submit" endpoint of the "vulnerabilities" service\.

## type SubmitRequestBody

SubmitRequestBody is the type of the "vulnerabilities" service "submit" endpoint HTTP request body\.

```go
type SubmitRequestBody struct {
    // Title of the vulnerability
    Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
    // Description of the vulnerability
    Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
    // If the vulnerability is exploitable
    Exploitable *bool `form:"exploitable,omitempty" json:"exploitable,omitempty" xml:"exploitable,omitempty"`
    // Severity score of the vulnerability
    CvssScore *float32 `form:"cvss_score,omitempty" json:"cvss_score,omitempty" xml:"cvss_score,omitempty"`
    // If the vulnerability is patchable
    IsPatchable *bool `form:"is_patchable,omitempty" json:"is_patchable,omitempty" xml:"is_patchable,omitempty"`
    // If the vulnerability is upgradeable
    IsUpgradeable *bool `form:"is_upgradeable,omitempty" json:"is_upgradeable,omitempty" xml:"is_upgradeable,omitempty"`
}
```

## type VulnerabilityResponse

VulnerabilityResponse is used to define fields on response body types\.

```go
type VulnerabilityResponse struct {
    // Unique ID of the vulnerability
    ID  uint64 `form:"id" json:"id" xml:"id"`
    // Title of the vulnerability
    Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
    // Description of the vulnerability
    Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
    // If the vulnerability is exploitable
    Exploitable *bool `form:"exploitable,omitempty" json:"exploitable,omitempty" xml:"exploitable,omitempty"`
    // Severity score of the vulnerability
    CvssScore *float32 `form:"cvss_score,omitempty" json:"cvss_score,omitempty" xml:"cvss_score,omitempty"`
    // If the vulnerability is patchable
    IsPatchable *bool `form:"is_patchable,omitempty" json:"is_patchable,omitempty" xml:"is_patchable,omitempty"`
    // If the vulnerability is upgradeable
    IsUpgradeable *bool `form:"is_upgradeable,omitempty" json:"is_upgradeable,omitempty" xml:"is_upgradeable,omitempty"`
}
```



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)