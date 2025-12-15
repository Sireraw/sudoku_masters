# Go Sudoku Solver API
**Quick Overview**
This project is a high-performance REST API built in Go (Golang) designed to solve standard 9x9 Sudoku puzzles. It uses the backtracking algorithm and leverages Go's built-in net/http package to provide a simple, robust network service.
The API exposes a single endpoint that accepts an unsolved Sudoku board (JSON payload where 0 is an empty cell) and returns the solved board.

**Core Technology Stack**
- Language: Go (Golang)
- Networking: net/http (Standard Library)
- Data Handling: encoding/json
- Algorithm: Backtracking

# Setup and Installation
**Prerequisites**
Ensure you have **Go 1.20 or newer** installed. Verify with:
```bash
go version
```

**Installation Steps**
Initialize Project:

```bash
mkdir go-sudoku-api
cd go-sudoku-api
go mod init go-sudoku-api
```
**Code Placement:** Place the main.go file (containing the solver and API logic) into the go-sudoku-api directory.

# Running & Usage
**Running the Server**
Start the API from the command line:
```bash
go run main.go
```

The server will start on port 8080: http://localhost:8080/solve

# API Endpoint
```bash
| Method | Path | Description |
POST
/solve
Submits a JSON payload containing the board array for solving.
```

## Example Request Payload Structure:
```bash
{ "board": [[5, 3, 0, ...], [...]] }
```

## Architectural Takeaways (Go vs. Python)
This project demonstrates Go's strengths compared to typical Python API setups:
- **Built-in Server:** Go requires no external frameworks (like Flask/FastAPI) for robust API functionality, relying entirely on the standard library.
- **Concurrency:** The net/http server automatically handles incoming requests using fast, lightweight Goroutines, ensuring the API remains non-blocking even when the solver is busy.
- **Type Safety:** Input and output are strictly managed using Go structs with JSON tags, providing compile-time assurance that is often handled dynamically in Python.

