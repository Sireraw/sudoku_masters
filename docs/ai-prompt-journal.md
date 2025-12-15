## AI Prompt Journal

This document serves as a detailed record of the generative AI prompts used during the learning process, the resulting code and architectural decisions, and the common issues encountered while building the Go Sudoku Solver API.

## AI Prompt Journal

This journal documents the prompts used to learn Go and structure the API, highlighting how the AI accelerated the learning and development process.

| Prompt | Summary of AI Response | Evaluation/Reflection |
|---------|-----------------------|-----------------------|
| "Write a simple Go program that solves a 9x9 sudoku game using backtracking and prints the result." | Provided the core solveSudoku, isSafe, and printBoard functions. (The base for this project). | **High Productivity.** Solved the core algorithm problem instantly. Showed correct use of array pointers (*[N][N]int) for in-place modification, a key Go concept. |
| "Convert the Sudoku solver Go code into a RESTful API using the net/http package. It should accept a JSON body with the board and return the solved board as JSON." | Generated the solveHandler function, demonstrated how to use json.NewDecoder and json.NewEncoder, and defined the necessary struct types (SudokuRequest, SudokuResponse). | **Excellent Clarity.** Clearly demonstrated how Go uses structs and "tags" for JSON marshaling/unmarshaling, which is crucial for API building in Go. |
| "In Go's net/http package, how do I handle CORS headers (Access-Control-Allow-Origin) for a POST request?" | Provided the necessary w.Header().Set() commands for CORS and explained the need to handle the http.MethodOptions (pre-flight request). | **Targeted Solution.** Saved significant time debugging browser-based testing issues by providing the standard CORS pattern immediately. |
| "Compare Go's approach to API building using its standard library vs. Python's Flask or FastAPI." | Detailed the difference between standard library vs. framework, compiled vs. interpreted, and Go's performance edge in concurrent I/O. | **Learning Reflection.** Helped frame the core architectural differences between Go and Python for the final documentation. |

# Common Issues & Fixes

| Issue | Description | Resolution |
|-------|-------------|-------------|
| Error: json: cannot unmarshal array into Go value of type main.SudokuRequest | The API client was sending the board data incorrectly, often without the encapsulating JSON object ({"board": [...]}). | **Fix:** Ensure the incoming JSON body exactly matches the SudokuRequest struct, including the "board" key. |
| API returns 405 Method Not Allowed | The client was sending an OPTIONS request (CORS pre-flight) which the handler wasn't explicitly set up to handle, or the wrong method was used. | **Fix:** Added the if r.Method == http.MethodOptions block in solveHandler and ensured the client uses POST. |
| Solver gets stuck/returns wrong answer | A subtle logic error, usually in the isValid checks, particularly the 3x3 box constraint. | **Fix:** Double-check the pointer math: startRow := row - row%3 and startCol := col - col%3 are correct for finding the top-left corner of the box. |
| Server only allows one connection at a time (Slow) | (Self-Correction) If you had written blocking code inside solveSudoku, it would block the server. Go's net/http package automatically handles each incoming request in its own Goroutine, which is why the server remains highly responsive even during complex computations. | **Fix:** (No fix needed, it's a Go feature!) The standard library's design prevents this issue, demonstrating Go's robust concurrency model. |

# References

1) Official Go Documentation: https://go.dev/doc/
2) Go Standard Library (net/http): https://pkg.go.dev/net/http (The primary tool for API creation).
3) Go Standard Library (encoding/json): https://pkg.go.dev/encoding/json (Essential for handling API data).
4) A Tour of Go: https://go.dev/tour/ (Excellent interactive tutorial for beginners).