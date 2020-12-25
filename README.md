# Citadel 🏰

## Table of contents

1. [Installation](#installation)
1. [About](#about)
1. [Vulnerability Service](#vulnerability-service)
1. [Incidents Service](#security-incident-service)

## Installation

> Clone the repository and navigate inside
```shell script
gh repo clone brittonhayes/citadel && cd citadel
```

> Download the vendor dependencies
```shell script
go mod download
```

> Navigate to the cmd directory, build, and run citadel
```shell script
cd ./cmd/citadel
go build
./citadel
```

## About

Citadel is a compilation of Security Operations (SecOps) microservices built with goa and go-kit. Each service functions independently, 
allowing you to use all of them or just the service that fits your needs. The goal of the project is to provide a baseline set of services 
that a team could easily store, query, and interact with SecOps resources. 

[Goa](https://github.com/goadesign/goa) allows this project to be very flexible to design changes. All the code and REST API documentation is generated based on the [design](./design/design.go) specification. 
The flexibility of this comes into play in the event that a resource or service needs to be expanded upon, added, or removed without untangling anything. 

### What's inside

Citadel includes the following services:

1. [Vulnerability Service](./gen/vulnerabilities)
1. [Security Incident Service](./gen/incidents)

#### Vulnerability Service

```yaml
Transports: 
    - http
Models: 
    - Vulnerability
```

#### Security Incident Service

```yaml
Transports: 
    - http
Models: 
    - Incident
```

---

> Built with [go-kit](https://github.com/go-kit/kit) and [gorm](https://gorm.io/)

> Social Icon by [Freepik](https://www.flaticon.com/authors/freepik) from [Flaticon](https://www.flaticon.com/)
