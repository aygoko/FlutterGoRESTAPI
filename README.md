# FlutterGoRESTAPI

A RESTful API built with Go for user management, designed to work with a Flutter frontend. This project includes user CRUD operations, Swagger documentation, and Docker support.

---

## Table of Contents

1. [Features](#features)
2. [Prerequisites](#prerequisites)
3. [Project Structure](#project-structure)
4. [Setup and Run](#setup-and-run)
5. [Testing the API](#testing-the-api)
6. [Swagger Documentation](#swagger-documentation)
7. [Docker Setup](#docker-setup)
8. [Contributing](#contributing)
9. [License](#license)

---

## Features

- **User Management**: Create, read, update, and delete users.
- **Swagger Docs**: Auto-generated API documentation.
- **Docker Support**: Containerized development and deployment.
- **Secure Passwords**: Passwords are hashed using `bcrypt`.
- **Clean Architecture**: Separation of concerns via domain, repository, and service layers.

---

## Prerequisites

- [Go](https://golang.org/) 1.24+
- [Flutter](https://flutter.dev/) (for the frontend)
- [Docker](https://www.docker.com/) (optional but recommended)
- [Postman](https://www.postman.com/) or `curl` for API testing

---

## Project Structure

```
your-project/
├── cmd/                # Entry point for the Go server
├── domain/             # Core business entities and interfaces
├── repository/         # Data storage implementations (e.g., RAM storage)
├── usecases/service/   # Business logic layer
├── go-api/             # HTTP handlers and routes
├── FlutterGoRESTAPI/   # Flutter frontend (if applicable)
├── Dockerfile          # Docker configuration
└── docker-compose.yml  # Docker service configuration
```

---

## Setup and Run

### 1. Clone the Repository

```bash
git clone https://github.com/aygoko/FlutterGoRESTAPI.git
cd FlutterGoRESTAPI
```

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Run the Server Locally

```bash
go run cmd/main.go
```

### 4. Access the API

- **Localhost**: `http://localhost:8080/api/users`
- **Swagger Docs**: [http://localhost:8080/swagger](http://localhost:8080/swagger) (requires Swagger UI setup)

---

## Testing the API

### Using `curl` (PowerShell Example)

```powershell
# Create a user
Invoke-WebRequest `
  -Uri "http://localhost:8080/api/users" `
  -Method POST `
  -Headers @{"Content-Type" = "application/json"} `
  -Body '{"login":"test", "email":"test@example.com", "password":"securepassword"}' `
  -UseBasicParsing

# Get all users
Invoke-WebRequest `
  -Uri "http://localhost:8080/api/users" `
  -Method GET `
  -UseBasicParsing

# Delete a user
Invoke-WebRequest `
  -Uri "http://localhost:8080/api/users/test_login" `
  -Method DELETE `
  -UseBasicParsing
```

---

## Swagger Documentation

1. **Generate Swagger Docs**:
   ```bash
   swag init  # Requires Swag installed (go get github.com/swaggo/swag/cmd/swag)
   ```

2. **Run Swagger UI**:
   ```bash
   swag serve
   ```
   Visit `http://localhost:8080/swagger` to explore endpoints.

---

## Docker Setup

### 1. Build the Docker Image

```bash
docker-compose up --build
```

### 2. Run the Docker Container

```bash
docker-compose up
```

### 3. Stop the Container

```bash
docker-compose down
```

---

## Contributing

1. **Report Issues**: Use GitHub Issues for bugs or feature requests.
2. **Pull Requests**: Fork the repo, make changes, and submit a PR.
3. **Code Style**: Follow Go conventions (e.g., `go fmt`, `go vet`).

---

## License

This project is licensed under the **MIT License**. See the [LICENSE](LICENSE) file for details.

---

## Contact

- **Author**: [Igor Kopytov](https://yourwebsite.com)
- **GitHub**: [aygoko](https://github.com/aygoko)

---

### Final Notes

- **Environment Variables**: Add environment-specific configurations (e.g., database settings) in `.env` files.
- **Production Ready**: Use the Docker image for deployment.
- **Security**: Add authentication (e.g., JWT) before deploying to production.

---
