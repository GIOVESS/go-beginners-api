# Go REST API Toolkit — Documentation

**Prompt-Powered Kickstart: Beginner's Toolkit for Go REST API**

Submitted by: Giovanni  
Repository: [https://github.com/GIOVESS/go-beginners-api](https://github.com/GIOVESS/go-beginners-api.git)

---

## 1. Overview of the Chosen Technology

### 1.1 Go (Golang)

**Go** is a statically typed, compiled programming language designed at Google. It emphasizes simplicity, readability, and efficient compilation.

| Feature | Description |
|--------|-------------|
| **Type** | Statically typed, compiled |
| **Use cases** | Backend services, REST APIs, CLI tools, microservices, cloud infrastructure |
| **Real-world examples** | Docker, Kubernetes, Prometheus, Terraform |
| **Strengths** | Fast compile times, built-in concurrency (goroutines), simple syntax, strong standard library |

### 1.2 Gin Web Framework

**Gin** is a high-performance HTTP web framework for Go. It provides routing, middleware, JSON validation, and error handling out of the box.

| Feature | Description |
|--------|-------------|
| **Purpose** | Build REST APIs and web services quickly |
| **Key capabilities** | Routing, middleware (logger, recovery), JSON binding, parameter validation |
| **Performance** | Very fast; minimal allocations |
| **Documentation** | [gin-gonic.com](https://gin-gonic.com/) |

### 1.3 Project Objective

This toolkit helps new developers:

- Create a minimal "Hello World" REST API in Go
- Understand Go modules (`go.mod`, `go.sum`)
- Learn routing and JSON responses
- Handle common setup and dependency issues

---

## 2. Setup Instructions

### 2.1 System Requirements

- **OS:** Linux, macOS, or Windows
- **Editor:** VS Code (or GoLand)
- **Go:** Installed and on your PATH — [go.dev/doc/install](https://go.dev/doc/install)
- **Dependencies:** Gin (listed in `go.mod`)

### 2.2 Install Go

1. Download the installer from [go.dev/dl](https://go.dev/dl/)
2. Run the installer and follow the prompts
3. Ensure Go is added to your system **PATH**

**Windows (custom install location):**  
If Go is installed in a non-standard path (e.g. `E:\Program Files\Go`):

- Add `E:\Program Files\Go\bin` to your user Environment Variables → Path, **or**
- Use the provided `run.ps1` script, which adds this path for the current session

### 2.3 Clone and Run the Project

```bash
# Clone the repository
git clone https://github.com/GIOVESS/go-beginners-api.git
cd go-beginners-api
```

**Option A — PowerShell script (Windows recommended)**

```powershell
.\run.ps1
```

This script:

- Adds Go to PATH (if needed)
- Runs `go mod tidy` to fetch dependencies
- Starts the API server

**Option B — Direct Go commands**

```bash
go mod tidy
go run main.go
```

The API runs at **http://localhost:8080**.

---

## 3. Minimal Working Example

### 3.1 Hello World API

The core example registers a single GET route `/hello` that returns a JSON greeting.

**main.go (minimal core)**

```go
package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // Hello World endpoint
    r.GET("/hello", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Hello, I'm Giovanni!",
        })
    })

    r.Run() // Runs on localhost:8080
}
```

### 3.2 Quick API Call

**Request:**

```
GET http://localhost:8080/hello
```

**Expected Response (200 OK):**

```json
{
  "message": "Hello, I'm Giovanni!"
}
```

### 3.3 Extended Endpoints (Joke API & Mini Chatbot)

The full codebase includes themed endpoints:

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/hello` | GET | Greeting message |
| `/joke` | GET | Random developer joke (optional `?topic=go`, `programming`, `api`) |
| `/chat?q=<message>` | GET | Minimal chatbot with canned responses |

**Example calls:**

- `GET http://localhost:8080/joke` → Returns a random joke in JSON
- `GET http://localhost:8080/chat?q=hello` → Returns a chatbot reply

---

## 4. AI Prompts Used & Learning Reflections

*Documentation reference: AI-assisted development and prompt engineering.*

| Prompt | AI Response Summary | Reflection |
|--------|---------------------|------------|
| "Show me how to create a minimal REST API in Go returning JSON." | Scaffolded project structure, Gin routing, and JSON response example | Very helpful for learning API basics |
| "Explain gin.Default() and routing in Go." | Explained middleware, default logger, and recovery | Improved understanding of Gin conventions |
| "Common errors with Go modules and solutions." | Suggested `go mod tidy` and dependency fixes | Saved troubleshooting time |
| "How to improve beginner-friendly Hello World API." | Suggested adding comments, error handling, and testing | Useful for creating clean, readable code |

---

## 5. Common Errors & How to Resolve Them

| Issue | Solution |
|-------|----------|
| `go` not recognized | Install Go and add to PATH; on Windows use `run.ps1` or add your Go `bin` folder (e.g. `E:\Program Files\Go\bin`) to PATH |
| missing go.sum entry / cannot find module | Run `go mod tidy` to download dependencies and create or update `go.sum` |
| Port 8080 already in use | Change to `r.Run(":8081")` in code, or stop the process using port 8080 |
| Gin version unknown (e.g. v1.9.3) | Use a valid version in `go.mod`, e.g. `github.com/gin-gonic/gin v1.10.0` |

---

## 6. Reference Resources

- **Go Official Documentation:** [go.dev/doc/](https://go.dev/doc/)
- **Go Installation:** [go.dev/doc/install](https://go.dev/doc/install)
- **Gin Web Framework:** [gin-gonic.com](https://gin-gonic.com/)
- **Gin GitHub:** [github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)
- Go tutorials and YouTube guides for REST APIs

---

## 7. Working Codebase

**Repository:**  
[https://github.com/GIOVESS/go-beginners-api](https://github.com/GIOVESS/go-beginners-api.git)

**Clone command:**

```bash
git clone https://github.com/GIOVESS/go-beginners-api.git
```

**Project structure:**

```
go-beginners-api/
├── main.go      # API server and endpoints
├── go.mod       # Go module definition
├── go.sum       # Dependency checksums
├── run.ps1      # Windows run script (adds PATH, runs server)
├── README.md    # Project readme
└── docs/        # Documentation (this file)
```

---

*End of Toolkit Document*
