# Hospital System API

A modern, scalable hospital management system built with Go, focusing on Clean Architecture, MVC patterns, and Object-Oriented Programming (OOP) principles.

## 🏗️ Project Structure

The project follows a **Clean Architecture** approach, ensuring separation of concerns and maintainability.

```text
hospital-system/
├── cmd/                        # Entry point of the application
│   ├── main.go                 # Application bootstrap
│   └── config.yaml             # App configuration file (Managed via Livecode)
├── config/                     # Configuration management (OOP-based, Livecode ready)
│   ├── config.go
│   ├── database.go
│   └── server.go
├── internal/
│   ├── driver/                 # Database drivers and connections
│   ├── middleware/             # HTTP middlewares
│   │   ├── cors.go             # CORS configuration
│   │   └── security.go         # Security headers and logic
│   ├── modules/                # Domain-driven modules (The Core)
│   │   └── health/             # Health Check Module
│   │       ├── controller/     # C: Controller (HTTP Handlers)
│   │       ├── entity/         # M: Model (Domain Entities)
│   │       └── service/        # M: Model (Business Logic Layer)
│   └── server/                 # Server setup and route registration
│       ├── gin.go              # Gin engine setup
│       ├── routes.go           # API route definitions
│       └── server.go           # Server startup logic
├── Dockerfile                  # Containerization
└── docker-compose.yml          # Local development environment
```

## 🛠️ Architecture & Design Patterns

### Clean Architecture (MVC)
We use a domain-driven MVC pattern within each module:
- **Model (M)**: Split into `entity` (data structures), `service` (business logic), and `repository` (data persistence).
- **View (V)**: In this API-first approach, the "View" is represented by JSON responses and DTOs.
- **Controller (C)**: Handlers that manage HTTP requests and orchestrate service calls.

### Object-Oriented Programming (OOP)
This project heavily utilizes OOP principles in Go:
- **Encapsulation**: Business logic is encapsulated within service and repository structs.
- **Interfaces**: Used to define contracts between layers, allowing for easy mocking and testing.
- **Dependency Injection (DI)**: Components are injected via constructors (e.g., `NewHealthController(service)`), making the system loosely coupled and highly testable.

## 🤖 AI-Assisted Development

This project is developed in collaboration with AI (Antigravity). We utilize AI for:
- **Architecture Design**: Ensuring Clean Architecture and OOP principles are strictly followed.
- **Livecode Integration**: Some parts of the structure are optimized for "Livecode" workflows, allowing for rapid prototyping and real-time code generation.
- **Code Refactoring**: Constant optimization of code for readability and performance.
- **Documentation**: Keeping the README and other docs up-to-date with structural changes.

## 🚀 Getting Started

### Prerequisites
- Go 1.22+
- Docker & Docker Compose

### Running the Application

1.  **Clone the repository**
    ```bash
    git clone <repository-url>
    cd hospital-system
    ```

2.  **Start with Docker (Database & App)**
    ```bash
    docker-compose up -d
    ```

3.  **Run locally**
    ```bash
    go run cmd/main.go
    ```

### API Endpoints
- **Health Check**: `GET /health` - Returns the current status of the API.

---
Created with ❤️ for High-Quality Go Development.
