### **f25/services/README.md**
# Backend Microservices

This folder contains all backend microservices for the **f25** project.  
We are building services using **Go** and following a **microservices architecture**.

## Structure
Each service is in its own directory:

```
services/
    service-name/
    main.go
    README.md
...
```

## Development Guidelines
- Each service is independent and should have its own `go.mod` file.
- Follow REST API best practices.
- Use environment variables for configuration.
- Keep business logic separate from HTTP handlers.

## Current Services
- _To be added_

## Getting Started
To run a service locally:
```bash
cd services/<service-name>
go run main.go
