# Citadel 🏰

[![Go Reference](https://pkg.go.dev/badge/github.com/brittonhayes/citadel.svg)](https://pkg.go.dev/github.com/brittonhayes/citadel)
[![Go Report Card](https://goreportcard.com/badge/github.com/brittonhayes/citadel)](https://goreportcard.com/report/github.com/brittonhayes/citadel)
<img width="64px" src="https://validator.swagger.io/validator/?url=https%3A%2F%2Fraw.githubusercontent.com%2Fbrittonhayes%2Fcitadel%2Fmain%2Fgen%2Fhttp%2Fopenapi3.json">

Citadel is a compilation of Security Operations (SecOps) microservices built
with [goa](https://github.com/goadesign/goa)

## Usage ✨

The citadel service includes a server as well as a command line interface. Let's get started using them together.

### Taskfile

This project uses [Taskfiles](https://taskfile.dev) as an alternative to Makefiles. To list the available commands run:

```shell
❯ task -l
task: Available tasks for this project:
* build: 		Build the project binaries
* build:distros: 	Build multiple distributions with goreleaser
* gen: 			Generate updated project files
* server: 		Run the citadel server
```

### Server

```shell
# Install the server
go get github.com/brittonhayes/citadel/cmd/citadel

# Start the server
go run citadel
```

### Client

```shell
# Install the CLI
go get github.com/brittonhayes/citadel/cmd/citadel-cli

# Run the CLI
go run citadel-cli
```

## Documentation

- [Open API 3 - YAML](./gen/http/openapi3.yaml)
- [Open API 3 - JSON](./gen/http/openapi3.json)

## Development

### GRPC Dependencies

```shell
# Install protobuf
brew install protobuf

# Install protobuf go plugin
go get -u github.com/golang/protobuf/protoc-gen-go
```

### Contributing

Want to contribute to Citadel? Feel free!

Any development to citadel should be done to `design files` and `business logic` files. Everything else is
auto-generated, so if you make changes to these files they will disappear.

### Project Structure

| Sections | Description | Type |
|------| ----------- | --------- |
| [incidents.go](incidents.go) | Business logic for the Security Incidents service. | `Business Logic` |
|[vulnerabilities.go](vulnerabilities.go) | Business logic for the Security Vulnerabilities service. |` Business Logic` |
|[design/design.go](design/design.go) | The primary service design  | `API Design` |
|[design/incident.go](design/incident.go) | Security incident service API design  | `API Design` |
|[design/vulnerability.go](design/vulnerability.go) | Security vulnerability service API design | `API design` |

---

## References

### Citadel

Security Operations Micro-Services

> Citadel is a compilation of Security Operations (SecOps) microservices built with goa. 
> Each service functions independently, allowing you to use all of them or just the service that fits your needs. 
> The goal of the project is to provide a baseline set of services that a team could easily store, query, and interact with SecOps resources.

### Goa

Design First Services

> [Goa](https://github.com/goadesign/goa) allows this project to be very flexible to design changes. 
> All the code and REST API documentation is generated based on the [design](./design/design.go) specification. 
> The flexibility of this comes into play in the event that a resource or service needs to be expanded upon, added, or removed without untangling anything.

---

> Social Icon by [Freepik](https://www.flaticon.com/authors/freepik) from [Flaticon](https://www.flaticon.com/)
