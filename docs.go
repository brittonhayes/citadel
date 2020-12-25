// citadel is a family of Security Operations (SecOps) microservices
//
// Introduction
//
// In Citadel, each service functions independently, allowing you to use all of
// them or just the service that fits your needs. The goal of the project
// is to provide a baseline set of services that a team could easily use to
// store, query, and interact with Security Operations resources.
//
// API Documentation
//
// REST API Documentation can be found in the following places:
//
// 	REST.md
//	gen/http/openapi3.yaml
//	gen/http/openapi3.json
//
// Installation
//
// Download the project, install dependencies, and generate updated docs and services
//
//  go get github.com/brittonhayes/citadel
//  go mod download
//  go generate
//
// Usage (Server)
//
// Navigate to the cmd/citadel directory to start up the server
//
//  cd cmd/citadel
//  go build
//  ./citadel
//
// Usage (CLI)
//
// Navigate to the cmd/citadel-cli directory to use the command line interface
//
//  cd cmd/citadel-cli
//  go build
//  ./citadel-cli --help
//
package citadel
