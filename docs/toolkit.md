# Beginner’s Toolkit: Go (Golang) - Sudoku Solver API

## 1. Title & Objective
**Technology Chosen:** Go (Golang)
**Objective:** To learn Go fundamentals, focusing on its architecture for concurrent and network-based applications, by building a simple, high-performance **Sudoku Solver REST API** using the backtracking algorithm.
**End Goal:** A runnable Go server that accepts an unsolved 9x9 Sudoku grid via a POST request (JSON) and returns the solved grid (JSON).

## 2. Quick Summary of the Technology
| **Feature** | **Description** |
|---------|-------------|
| **Name**| Go (often called Golang) |
| **Type**| Statically typed, compiled programming language |
| **Origin** | Developed by Google (Robert Griesemer, Rob Pike, and Ken Thompson) |
| **Concurrency** | Built-in, using Goroutines (lightweight threads) and Channels (for communication)
| **Use Cases** | Cloud services, networking, command-line tools, microservices, and high-performance APIs. |
| **Real-World Example**| **Docker** (containerization platform), **Kubernetes** (container orchestration), and **Prometheus** (monitoring system) are all written in Go. |

**Go Architecture vs. Python:**
| **Feature** | **Go (Golang)** | **Python** |
|---------|-------------|--------|
| **Execution** | **Compiled.** Generates a single binary file. Faster startup and runtime performance. | **Interpreted.** Requires a runtime environment (interpreter). |
| **Typing** | **Statically Typed.** Variables must have a defined type at compile time. Catches errors earlier. | **Dynamically Typed.** Types are checked at runtime. |
| **Concurrency** | **Goroutines.** Extremely lightweight and efficient concurrency model built into the language. | **Threading/Asyncio.** Uses OS threads or event loops; generally less efficient for high-concurrency I/O. |
| **Networking** | **Standard Library.** High-quality, robust net/http package handles APIs without external frameworks. | **External Frameworks** (Flask, Django, FastAPI) needed for robust API development. |
| **Error Handling** | **Explicit.** Uses multi-value returns (result, error). Developers must check errors explicitly. | **Implicit.** Uses try...except blocks. |

## 3. System Requirements
- **Operating System:** Windows, macOS, or Linux (any modern distribution).
- **Go Compiler/Runtime:** Go 1.20 or newer.
- **Text Editor:** VS Code (recommended) or any other IDE with Go extension.
- **API Client:** A tool like curl or Postman (or an in-browser console) to test the API endpoint.

## 4. Installation & Setup Instructions
 1. **Install Go:**
- Navigate to the official Go website: https://go.dev/doc/install
- Download the appropriate installer for your OS and follow the instructions.
- **Verify Installation:** Open your terminal/command prompt and run: go version
Expected output: go version go1.x.x <os>/<arch>

  2. **Initialize Project:**
- Create a new directory for your project and navigate into it.
- Initialize a Go module (this manages dependencies, though we use only standard library here):
```bash
mkdir go-sudoku-api
cd go-sudoku-api
go mod init go-sudoku-api
# Creates go.mod file
```

  3. **Create Code File:**
- Create the file main.go and paste the API code from Section 5.

## 5. Minimal Working Example (The Go API)
The provided main.go file implements the entire solver and API server.

**Code Explanation (Brief)**
The core architecture is:

  1. **main():** Sets up a router (http.HandleFunc) to map the /solve path to the solveHandler function and starts the server on port 8080 (http.ListenAndServe).
  2. **solveHandler():** Reads the incoming request body, decodes the JSON into a SudokuRequest struct, calls the solveSudoku function, and sends the result back as a SudokuResponse JSON.
  3. **solveSudoku():** The recursive backtracking implementation that mutates the board pointer until a solution is found (returns true) or all paths are exhausted (returns false).

**Running the Server**
Execute the Go program:
go run main.go
# Output: Sudoku Solver API is running on http://localhost:8080/solve


**Testing the API (Expected Output)**
Use curl in a separate terminal window to send a sample unsolved puzzle.

**Request Body (JSON):**
```bash
{
    "board": [
        [5, 3, 0, 0, 7, 0, 0, 0, 0],
        [6, 0, 0, 1, 9, 5, 0, 0, 0],
        [0, 9, 8, 0, 0, 0, 0, 6, 0],
        [8, 0, 0, 0, 6, 0, 0, 0, 3],
        [4, 0, 0, 8, 0, 3, 0, 0, 1],
        [7, 0, 0, 0, 2, 0, 0, 0, 6],
        [0, 6, 0, 0, 0, 0, 2, 8, 0],
        [0, 0, 0, 4, 1, 9, 0, 0, 5],
        [0, 0, 0, 0, 8, 0, 0, 7, 9]
    ]
}
```

**Curl Command:**
```bash
curl -X POST http://localhost:8080/solve \
     -H "Content-Type: application/json" \
     -d '{ "board": [[5, 3, 0, 0, 7, 0, 0, 0, 0], [6, 0, 0, 1, 9, 5, 0, 0, 0], [0, 9, 8, 0, 0, 0, 0, 6, 0], [8, 0, 0, 0, 6, 0, 0, 0, 3], [4, 0, 0, 8, 0, 3, 0, 0, 1], [7, 0, 0, 0, 2, 0, 0, 0, 6], [0, 6, 0, 0, 0, 0, 2, 8, 0], [0, 0, 0, 4, 1, 9, 0, 0, 5], [0, 0, 0, 0, 8, 0, 0, 7, 9]] }'
```

**Expected Output (JSON Response):**
```bash
{
    "solved": true,
    "solution": [
        [5, 3, 4, 6, 7, 8, 9, 1, 2],
        [6, 7, 2, 1, 9, 5, 3, 4, 8],
        [1, 9, 8, 3, 4, 2, 5, 6, 7],
        [8, 5, 9, 7, 6, 1, 4, 2, 3],
        [4, 2, 6, 8, 5, 3, 7, 9, 1],
        [7, 1, 3, 9, 2, 4, 8, 5, 6],
        [9, 6, 1, 5, 3, 7, 2, 8, 4],
        [2, 8, 7, 4, 1, 9, 6, 3, 5],
        [3, 4, 5, 2, 8, 6, 1, 7, 9]
    ],
    "message": "Sudoku puzzle solved successfully."
}
```

## 6. AI Prompt Journal- ( ai_prompt_journal.md)
| **Prompt** | **Summary of AI Response** | **Evaluation/Reflection** |
|------------|----------------------------|---------------------------|
| "Write a simple Go program that solves a 9x9 sudoku game using backtracking and prints the result." | Provided the core solveSudoku, isSafe, and printBoard functions. (The base for this project). | **High Productivity**. Solved the core algorithm problem instantly. Showed correct use of array pointers (*[N][N]int) for in-place modification. |
| "Convert the Sudoku solver Go code into a RESTful API using the net/http package. It should accept a JSON body with the board and return the solved board as JSON." | Generated the solveHandler function, demonstrated how to use json.NewDecoder and json.NewEncoder, and defined the necessary struct types (SudokuRequest, SudokuResponse). | **Excellent Clarity**. Clearly demonstrated how Go uses structs for JSON marshaling/unmarshaling, which is crucial for API building in Go.|
| "In Go's net/http package, how do I handle CORS headers (Access-Control-Allow-Origin) for a POST request?" | Provided the necessary w.Header().Set() commands for CORS and explained the need to handle the http.MethodOptions (pre-flight request). | **Targeted Solution.** Saved significant time debugging browser-based testing issues by providing the standard CORS pattern. |
| "Compare Go's approach to API building using its standard library vs. Python's Flask or FastAPI." | Detailed the difference between standard library vs. framework, compiled vs. interpreted, and Go's performance edge in concurrent I/O. | **Learning Reflection.** Helped frame the "Quick Summary of the Technology" section and solidify the differences in language architecture. |

## 7. Common Issues & Fixes
| **Issue** | **Description** | **Resolution** |
|-------|-------------|------------|
| **Error: json: cannot unmarshal array into Go value of type main.SudokuRequest** | The API client was sending the board data incorrectly, often without the encapsulating JSON object ({"board": [...]}). | **Fix:** Ensure the incoming JSON body exactly matches the SudokuRequest struct, including the "board" key. |
| **API returns 405 Method Not Allowed** | The client was sending an OPTIONS request (CORS pre-flight) which the handler wasn't explicitly set up to handle, or the wrong method was used. | **Fix:** Added the if r.Method == http.MethodOptions block in solveHandler and ensured the client uses POST. |
| **Solver gets stuck/returns wrong answer** | A subtle logic error, usually in the isValid checks, particularly the 3x3 box constraint. |  Fix: Double-check the pointer math: startRow := row - row%3 and startCol := col - col%3 are correct for finding the top-left corner of the box. |
| **Server only allows one connection at a time (Slow)** | (Self-Correction) If you had written blocking code inside solveSudoku, it would block the server. Go's net/http package automatically handles each incoming request in its own **Goroutine**, which is why the server remains highly responsive even during complex computations. | **Fix:** (No fix needed, it's a Go feature!) The standard library's design prevents this issue, demonstrating Go's robust concurrency model. |

## 8. References
- **Official Go Documentation:** https://go.dev/doc/
- **Go Standard Library (net/http):** https://pkg.go.dev/net/http (The primary tool for API creation).
- **Go Standard Library (encoding/json):** https://pkg.go.dev/encoding/json (Essential for handling API data).
- **A Tour of Go:** https://go.dev/tour/ (Excellent interactive tutorial for beginners).
