// Code generated by goa v3.2.6, DO NOT EDIT.
//
// incidents endpoints
//
// Command:
// $ goa gen citadel/design

package incidents

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints wraps the "incidents" service endpoints.
type Endpoints struct {
	Find    endpoint.Endpoint
	ListAll endpoint.Endpoint
}

// NewEndpoints wraps the methods of the "incidents" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Find:    NewFindEndpoint(s),
		ListAll: NewListAllEndpoint(s),
	}
}

// Use applies the given middleware to all the "incidents" service endpoints.
func (e *Endpoints) Use(m func(endpoint.Endpoint) endpoint.Endpoint) {
	e.Find = m(e.Find)
	e.ListAll = m(e.ListAll)
}

// NewFindEndpoint returns an endpoint function that calls the method "find" of
// service "incidents".
func NewFindEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*FindPayload)
		return s.Find(ctx, p)
	}
}

// NewListAllEndpoint returns an endpoint function that calls the method "list
// all" of service "incidents".
func NewListAllEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*LimitPayload)
		return s.ListAll(ctx, p)
	}
}
