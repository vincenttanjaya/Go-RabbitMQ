# Go-RabbitMQ

## Project Structure
```
.
├── app                 // Main applications for this project.
│   └── main.go
├── constant            // Constant for this project.
│   └── constant.go
├── http                // Http and routes for this project (include handler)
│   └── api.go
├── mock                // Mock files from repository
│   ├── rabbitmq        // Mock for rabbitmq repository (for unit test)
│       └── rabbitmq.go 
├── model               // Include all struct used in this service (include request and response)
│   └── model.go
├── repository/rabbitmq // RabbitMQ Repository
│   ├── consume.go      // Implementation for consume rabbitmq
│   ├── interface.go    // Interface for rabbitmq repository
│   └── publish.go      // Implementation for publish rabbitmq
├── service             // usecase for this project
│   ├── service_test.go // unit test for service func
│   └── service.go      // implementation and interface for servce
├── utils               // helper function for this project
│   ├── config.go       // helper to read config file
│   └── rabbitmq.go     // helper for goroutine consume and init
├── config.yaml         // configuration for this project
├── go.mod
├── go.sum
└── README.md
```
## How to run this project
```
1. go mod tidy / go mod vendor
2. go run app/main.go

this project will run on localhost:8080
the port can be customize by change the config value in the config.yaml

for the configuration for rabbitmq can change in the config.yaml
```