# Complete_microservices_with_go

🚀A learning project building while taking "Complete microservices with go" Udemy course

**Course:** [Complete microservices with go](https://www.udemy.com/course/complete-microservices-with-go)

```bash
go test -v *.go
```


```bash
minikube dashboard
```

```bash
tilt up
```

```bash
go run tools/create_service.go -name test
```
```
Directory structure created:

services/test-service/
├── cmd/                    # Application entry points
├── internal/              # Private application code
│   ├── domain/           # Business domain models and interfaces
│   │   └── test.go         # Core domain interfaces
│   ├── service/          # Business logic implementation
│   │   └── service.go    # Service implementations
│   └── infrastructure/   # External dependencies implementations (abstractions)
│       ├── events/       # Event handling (RabbitMQ)
│       ├── grpc/         # gRPC server handlers
│       └── repository/   # Data persistence
├── pkg/                  # Public packages
│   └── types/           # Shared types and models
└── README.md            # This file
```