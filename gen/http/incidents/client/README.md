<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# client

```go
import "citadel/gen/http/incidents/client"
```

## Index

- [func BuildFindPayload(incidentsFindID string) (*incidents.FindPayload, error)](<#func-buildfindpayload>)
- [func BuildListAllPayload(incidentsListAllLimit string) (*incidents.LimitPayload, error)](<#func-buildlistallpayload>)
- [func DecodeFindResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error)](<#func-decodefindresponse>)
- [func DecodeListAllResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error)](<#func-decodelistallresponse>)
- [func EncodeListAllRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error](<#func-encodelistallrequest>)
- [func FindIncidentsPath(id uint64) string](<#func-findincidentspath>)
- [func ListAllIncidentsPath() string](<#func-listallincidentspath>)
- [func NewFindIncidentOK(body *FindResponseBody) *incidents.Incident](<#func-newfindincidentok>)
- [func NewListAllIncidentOK(body []*IncidentResponse) []*incidents.Incident](<#func-newlistallincidentok>)
- [func ValidateFindResponseBody(body *FindResponseBody) (err error)](<#func-validatefindresponsebody>)
- [func ValidateIncidentResponse(body *IncidentResponse) (err error)](<#func-validateincidentresponse>)
- [type Client](<#type-client>)
  - [func NewClient(
    scheme string,
    host string,
    doer goahttp.Doer,
    enc func(*http.Request) goahttp.Encoder,
    dec func(*http.Response) goahttp.Decoder,
    restoreBody bool,
) *Client](<#func-newclient>)
  - [func (c *Client) BuildFindRequest(ctx context.Context, v interface{}) (*http.Request, error)](<#func-client-buildfindrequest>)
  - [func (c *Client) BuildListAllRequest(ctx context.Context, v interface{}) (*http.Request, error)](<#func-client-buildlistallrequest>)
  - [func (c *Client) Find() endpoint.Endpoint](<#func-client-find>)
  - [func (c *Client) ListAll() endpoint.Endpoint](<#func-client-listall>)
- [type FindResponseBody](<#type-findresponsebody>)
- [type IncidentResponse](<#type-incidentresponse>)
- [type ListAllResponseBody](<#type-listallresponsebody>)


## func BuildFindPayload

```go
func BuildFindPayload(incidentsFindID string) (*incidents.FindPayload, error)
```

BuildFindPayload builds the payload for the incidents find endpoint from CLI flags\.

## func BuildListAllPayload

```go
func BuildListAllPayload(incidentsListAllLimit string) (*incidents.LimitPayload, error)
```

BuildListAllPayload builds the payload for the incidents list all endpoint from CLI flags\.

## func DecodeFindResponse

```go
func DecodeFindResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error)
```

DecodeFindResponse returns a decoder for responses returned by the incidents find endpoint\. restoreBody controls whether the response body should be restored after having been read\.

## func DecodeListAllResponse

```go
func DecodeListAllResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error)
```

DecodeListAllResponse returns a decoder for responses returned by the incidents list all endpoint\. restoreBody controls whether the response body should be restored after having been read\.

## func EncodeListAllRequest

```go
func EncodeListAllRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error
```

EncodeListAllRequest returns an encoder for requests sent to the incidents list all server\.

## func FindIncidentsPath

```go
func FindIncidentsPath(id uint64) string
```

FindIncidentsPath returns the URL path to the incidents service find HTTP endpoint\.

## func ListAllIncidentsPath

```go
func ListAllIncidentsPath() string
```

ListAllIncidentsPath returns the URL path to the incidents service list all HTTP endpoint\.

## func NewFindIncidentOK

```go
func NewFindIncidentOK(body *FindResponseBody) *incidents.Incident
```

NewFindIncidentOK builds a "incidents" service "find" endpoint result from a HTTP "OK" response\.

## func NewListAllIncidentOK

```go
func NewListAllIncidentOK(body []*IncidentResponse) []*incidents.Incident
```

NewListAllIncidentOK builds a "incidents" service "list all" endpoint result from a HTTP "OK" response\.

## func ValidateFindResponseBody

```go
func ValidateFindResponseBody(body *FindResponseBody) (err error)
```

ValidateFindResponseBody runs the validations defined on FindResponseBody

## func ValidateIncidentResponse

```go
func ValidateIncidentResponse(body *IncidentResponse) (err error)
```

ValidateIncidentResponse runs the validations defined on IncidentResponse

## type Client

Client lists the incidents service endpoint HTTP clients\.

```go
type Client struct {
    // Find Doer is the HTTP client used to make requests to the find endpoint.
    FindDoer goahttp.Doer

    // ListAll Doer is the HTTP client used to make requests to the list all
    // endpoint.
    ListAllDoer goahttp.Doer

    // RestoreResponseBody controls whether the response bodies are reset after
    // decoding so they can be read again.
    RestoreResponseBody bool
    // contains filtered or unexported fields
}
```

### func NewClient

```go
func NewClient(
    scheme string,
    host string,
    doer goahttp.Doer,
    enc func(*http.Request) goahttp.Encoder,
    dec func(*http.Response) goahttp.Decoder,
    restoreBody bool,
) *Client
```

NewClient instantiates HTTP clients for all the incidents service servers\.

### func \(\*Client\) BuildFindRequest

```go
func (c *Client) BuildFindRequest(ctx context.Context, v interface{}) (*http.Request, error)
```

BuildFindRequest instantiates a HTTP request object with method and path set to call the "incidents" service "find" endpoint

### func \(\*Client\) BuildListAllRequest

```go
func (c *Client) BuildListAllRequest(ctx context.Context, v interface{}) (*http.Request, error)
```

BuildListAllRequest instantiates a HTTP request object with method and path set to call the "incidents" service "list all" endpoint

### func \(\*Client\) Find

```go
func (c *Client) Find() endpoint.Endpoint
```

Find returns an endpoint that makes HTTP requests to the incidents service find server\.

### func \(\*Client\) ListAll

```go
func (c *Client) ListAll() endpoint.Endpoint
```

ListAll returns an endpoint that makes HTTP requests to the incidents service list all server\.

## type FindResponseBody

FindResponseBody is the type of the "incidents" service "find" endpoint HTTP response body\.

```go
type FindResponseBody struct {
    // Unique ID of the incident
    ID  *uint64 `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
    // Date the incident occurred
    Date *string `form:"date,omitempty" json:"date,omitempty" xml:"date,omitempty"`
    // Date the incident occurred
    DateClosed *string `form:"date_closed,omitempty" json:"date_closed,omitempty" xml:"date_closed,omitempty"`
    // Permissions associated with incident
    Permissions *string `form:"Permissions,omitempty" json:"Permissions,omitempty" xml:"Permissions,omitempty"`
    // The severity of the incident
    Severity *int `form:"severity,omitempty" json:"severity,omitempty" xml:"severity,omitempty"`
    // The short title of the incident
    Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
    // The detailed description of the incident
    Summary *string `form:"summary,omitempty" json:"summary,omitempty" xml:"summary,omitempty"`
    // The scope of impact of this incident
    Scope *string `form:"scope,omitempty" json:"scope,omitempty" xml:"scope,omitempty"`
    // What group or individual caused the initial incident
    ResponsibleParty *string `form:"responsible_party,omitempty" json:"responsible_party,omitempty" xml:"responsible_party,omitempty"`
    // A list of the affected customers
    AffectedCustomers []string `form:"affected_customers,omitempty" json:"affected_customers,omitempty" xml:"affected_customers,omitempty"`
    // The original cause of the incident
    RootCause *string `form:"root_cause,omitempty" json:"root_cause,omitempty" xml:"root_cause,omitempty"`
    // The slack channel for incident discussions
    SlackChannel *string `form:"slack_channel,omitempty" json:"slack_channel,omitempty" xml:"slack_channel,omitempty"`
    // When the incident was submitted
    CreatedAt *string `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
    // When the incident was last updated
    UpdatedAt *string `form:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}
```

## type IncidentResponse

IncidentResponse is used to define fields on response body types\.

```go
type IncidentResponse struct {
    // Unique ID of the incident
    ID  *uint64 `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
    // Date the incident occurred
    Date *string `form:"date,omitempty" json:"date,omitempty" xml:"date,omitempty"`
    // Date the incident occurred
    DateClosed *string `form:"date_closed,omitempty" json:"date_closed,omitempty" xml:"date_closed,omitempty"`
    // Permissions associated with incident
    Permissions *string `form:"Permissions,omitempty" json:"Permissions,omitempty" xml:"Permissions,omitempty"`
    // The severity of the incident
    Severity *int `form:"severity,omitempty" json:"severity,omitempty" xml:"severity,omitempty"`
    // The short title of the incident
    Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
    // The detailed description of the incident
    Summary *string `form:"summary,omitempty" json:"summary,omitempty" xml:"summary,omitempty"`
    // The scope of impact of this incident
    Scope *string `form:"scope,omitempty" json:"scope,omitempty" xml:"scope,omitempty"`
    // What group or individual caused the initial incident
    ResponsibleParty *string `form:"responsible_party,omitempty" json:"responsible_party,omitempty" xml:"responsible_party,omitempty"`
    // A list of the affected customers
    AffectedCustomers []string `form:"affected_customers,omitempty" json:"affected_customers,omitempty" xml:"affected_customers,omitempty"`
    // The original cause of the incident
    RootCause *string `form:"root_cause,omitempty" json:"root_cause,omitempty" xml:"root_cause,omitempty"`
    // The slack channel for incident discussions
    SlackChannel *string `form:"slack_channel,omitempty" json:"slack_channel,omitempty" xml:"slack_channel,omitempty"`
    // When the incident was submitted
    CreatedAt *string `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
    // When the incident was last updated
    UpdatedAt *string `form:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}
```

## type ListAllResponseBody

ListAllResponseBody is the type of the "incidents" service "list all" endpoint HTTP response body\.

```go
type ListAllResponseBody []*IncidentResponse
```



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)