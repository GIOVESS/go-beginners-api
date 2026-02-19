# Prompt-Powered Kickstart: Beginner’s Toolkit for Go REST API

## Objective
This project is a beginner-friendly toolkit for learning **Go (Golang)** and building a minimal REST API.  
The goal is to help new developers quickly get started with Go, create a simple "Hello, World!" API, and understand key concepts such as modules, routing, and JSON responses.

---

## Quick Summary of Go
- Go is a **statically typed, compiled language** developed by Google.  
- Used for backend services, APIs, CLI tools, and microservices.  
- Real-world example: **Docker and Kubernetes are built in Go.**

---

## System Requirements
- OS: Linux, Mac, or Windows  
- Tools: VS Code (or GoLand)  
- Go installation: [Official Go Docs](https://golang.org/doc/install)  
- Packages: `gin` for REST API (optional, can use `net/http`)

---

## Installation & Setup

```bash
# Clone the repository
git clone https://github.com/GIOVESS/go-beginners-api.git
cd go-beginners-api

# Initialize Go module (if starting fresh)
go mod init beginners-api

# Install Gin framework
go get -u github.com/gin-gonic/gin
