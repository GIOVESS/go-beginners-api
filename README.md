# Prompt-Powered Kickstart: Beginner’s Toolkit for Go REST API

## Objective
This project is a beginner-friendly toolkit for learning **Go (Golang)** and building a minimal REST API.  
The goal is to help new developers quickly get started with Go, create a simple "Hello, I'm Giovanni!" API, and understand key concepts such as modules, routing, and JSON responses.

---

## Quick Summary of Go
- Go is a **statically typed, compiled language** developed by Google.  
- Used for backend services, APIs, CLI tools, and microservices.  
- Real-world example: **Docker and Kubernetes are built in Go.**

---

## System Requirements
- OS: Linux, Mac, or Windows  
- Tools: VS Code (or GoLand)  
- **Go installed** and on your PATH: [Official Go Docs](https://golang.org/doc/install)  
- Packages: `gin` for REST API (listed in `go.mod`)

---

## Installation & Setup

### 1. Install Go
- Install Go from [go.dev/dl](https://go.dev/dl/) and ensure the installer adds Go to your **PATH**.
- On Windows, if Go is installed in a custom location (e.g. `E:\Program Files\Go`), either:
  - Add `E:\Program Files\Go\bin` to your user PATH (Environment Variables), or  
  - Use `run.ps1` (see below), which adds that path for the current session.

### 2. Get the project

```bash
# Clone the repository
git clone https://github.com/GIOVESS/go-beginners-api.git
cd go-beginners-api
```

### 3. Run the API

**Option A – Using the PowerShell script (recommended on Windows)**  
`run.ps1` adds Go to PATH for the session (if needed), runs `go mod tidy`, then starts the API:

```powershell
.\run.ps1
```

**Option B – Using Go directly**  
If `go` is already on your PATH:

```bash
go mod tidy
go run main.go
```

The API runs at **http://localhost:8080**. Try: [http://localhost:8080/hello](http://localhost:8080/hello).

---

### 4. Minimal Working Example

This example starts an HTTP server with the Gin framework, registers a single GET route `/hello`, and responds with a JSON object. It illustrates a minimal REST API: one endpoint, one handler, JSON response.

**main.go**

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Hello Giovanni endpoint
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, I'm Giovanni!",
		})
	})

	var err = r.Run()
	if err != nil {
		return
	} // Runs on localhost:8080
}

```

**Expected output**

`GET http://localhost:8080/hello` →

```json
{ "message": "Hello, I'm Giovanni!" }
```

---

## AI Prompt Journal

| Prompt | AI Response Summary | Reflection |
|--------|---------------------|------------|
| "Show me how to create a minimal REST API in Go returning JSON." | Scaffolded project structure and routing example | Very helpful for learning API basics |
| "Explain gin.Default() and routing in Go." | Explained middleware, default logger, and recovery | Improved understanding of Gin conventions |
| "Common errors with Go modules and solutions." | Suggested go mod tidy and dependency fixes | Saved troubleshooting time |
| "How to improve beginner-friendly Hello World API." | Suggested adding comments, error handling, and testing | Useful for creating clean, readable code |

---

## Common Issues & Fixes

| Issue | Solution |
|-------|----------|
| `go` not recognized / cannot find module | Install Go and add to PATH; on Windows use `run.ps1` or add `E:\Program Files\Go\bin` (or your install path) to PATH |
| missing go.sum entry | Run `go mod tidy` to download dependencies and create/update `go.sum` |
| Port 8080 already in use | Change to `r.Run(":8081")` or free the port |
| Gin version unknown (e.g. v1.9.3) | Use a valid version in `go.mod`, e.g. `github.com/gin-gonic/gin v1.10.0` |

---

## References

- [Go Official Documentation](https://go.dev/doc/)
- [Gin Web Framework](https://gin-gonic.com/)
- Go tutorials and YouTube guides for REST APIs
