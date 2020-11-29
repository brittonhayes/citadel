# Citadel 🏰

## Table of contents

1. [Installation](#Installation)
1. [About](#About)
1. [Vulnerability Service](#Vulnerability Service)
1. [Risk Service](#Risk Service)
1. [Threat Intel Service](#Threat Intel Service)

## Installation

```shell script
# Clone the repository
git clone git@github.com:brittonhayes/citadel.git

# Start it up in Docker
docker-compose up --build
```

## About

Citadel is a compilation of Security Operations (SecOps) microservices built with go-kit. Each service functions independently, allowing you to use all of them or just the service that fits your needs. 
The goal of the project is to provide a baseline set of services that a team could easily deploy and have means of storing, querying, and interacting with SecOps resources. 

### What's inside

Citadel includes the following services:

1. [Vulnerability Service](./vulnerability)
2. [Risk Service](./risk)
3. [Threat Intel Service](#Threat intel service)

#### Vulnerability Service

```yaml
Status: In Development
Transports: 
    - grpc
    - http
Models: 
    - Vulnerability
    - Patch
Storage:
    - Sqlite
    - Postgres
    - MySql
```


#### Risk Service

```yaml
Status: Not Started
Transports: 
    - grpc
    - http
Models: 
    - Risk
    - Risk Adjustment Request
Storage:
    - Sqlite
    - Postgres
    - MySql
```

#### Threat Intel Service

```yaml
Status: Not Started
```

---

> Built with [go-kit](https://github.com/go-kit/kit) and [gorm](https://gorm.io/)

