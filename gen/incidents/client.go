// Code generated by goa v3.2.6, DO NOT EDIT.
//
// incidents client
//
// Command:
// $ goa gen github.com/brittonhayes/citadel/design

package incidents

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "incidents" service client.
type Client struct {
	FindEndpoint    goa.Endpoint
	ListAllEndpoint goa.Endpoint
}

// NewClient initializes a "incidents" service client given the endpoints.
func NewClient(find, listAll goa.Endpoint) *Client {
	return &Client{
		FindEndpoint:    find,
		ListAllEndpoint: listAll,
	}
}

// Find calls the "find" endpoint of the "incidents" service.
// Find may return the following errors:
//	- "no_match" (type NoMatch)
//	- error: internal error
func (c *Client) Find(ctx context.Context, p *FindPayload) (res *Incident, err error) {
	var ires interface{}
	ires, err = c.FindEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Incident), nil
}

// ListAll calls the "list all" endpoint of the "incidents" service.
// ListAll may return the following errors:
//	- "no_match" (type NoMatch)
//	- error: internal error
func (c *Client) ListAll(ctx context.Context, p *LimitPayload) (res []*Incident, err error) {
	var ires interface{}
	ires, err = c.ListAllEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.([]*Incident), nil
}
