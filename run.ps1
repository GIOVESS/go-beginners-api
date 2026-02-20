# Add Go to PATH for this session (Go installed at E:\Program Files\Go)
$env:Path = "E:\Program Files\Go\bin;" + $env:Path

# Download dependencies and update go.sum
go mod tidy

# Run the API
go run main.go
