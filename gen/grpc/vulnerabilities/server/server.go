// Code generated by goa v3.2.6, DO NOT EDIT.
//
// vulnerabilities gRPC server
//
// Command:
// $ goa gen github.com/brittonhayes/citadel/design

package server

import (
	"context"

	vulnerabilitiespb "github.com/brittonhayes/citadel/gen/grpc/vulnerabilities/pb"
	vulnerabilities "github.com/brittonhayes/citadel/gen/vulnerabilities"
	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
)

// Server implements the vulnerabilitiespb.VulnerabilitiesServer interface.
type Server struct {
	FindH goagrpc.UnaryHandler
	ListH goagrpc.UnaryHandler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the expr.
type ErrorNamer interface {
	ErrorName() string
}

// New instantiates the server struct with the vulnerabilities service
// endpoints.
func New(e *vulnerabilities.Endpoints, uh goagrpc.UnaryHandler) *Server {
	return &Server{
		FindH: NewFindHandler(e.Find, uh),
		ListH: NewListHandler(e.List, uh),
	}
}

// NewFindHandler creates a gRPC handler which serves the "vulnerabilities"
// service "find" endpoint.
func NewFindHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeFindRequest, EncodeFindResponse)
	}
	return h
}

// Find implements the "Find" method in vulnerabilitiespb.VulnerabilitiesServer
// interface.
func (s *Server) Find(ctx context.Context, message *vulnerabilitiespb.FindRequest) (*vulnerabilitiespb.FindResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "find")
	ctx = context.WithValue(ctx, goa.ServiceKey, "vulnerabilities")
	resp, err := s.FindH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*vulnerabilitiespb.FindResponse), nil
}

// NewListHandler creates a gRPC handler which serves the "vulnerabilities"
// service "list" endpoint.
func NewListHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeListRequest, EncodeListResponse)
	}
	return h
}

// List implements the "List" method in vulnerabilitiespb.VulnerabilitiesServer
// interface.
func (s *Server) List(ctx context.Context, message *vulnerabilitiespb.ListRequest) (*vulnerabilitiespb.ListResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "list")
	ctx = context.WithValue(ctx, goa.ServiceKey, "vulnerabilities")
	resp, err := s.ListH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*vulnerabilitiespb.ListResponse), nil
}
