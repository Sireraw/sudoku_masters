package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const N = 9

// --- Sudoku Algorithm ---

func isSafe(board *[N][N]int, row, col, num int) bool {
	for c := 0; c < N; c++ {
		if board[row][c] == num { return false }
	}
	for r := 0; r < N; r++ {
		if board[r][col] == num { return false }
	}
	sr, sc := row-row%3, col-col%3
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			if board[sr+r][sc+c] == num { return false }
		}
	}
	return true
}

func solveSudoku(board *[N][N]int) bool {
	row, col, found := -1, -1, false
	for r := 0; r < N && !found; r++ {
		for c := 0; c < N; c++ {
			if board[r][c] == 0 {
				row, col, found = r, c, true
				break
			}
		}
	}
	if !found { return true }
	for num := 1; num <= 9; num++ {
		if isSafe(board, row, col, num) {
			board[row][col] = num
			if solveSudoku(board) { return true }
			board[row][col] = 0
		}
	}
	return false
}

// --- API Communication Structures ---

type SolveRequest struct {
	Board [N][N]int `json:"board"`
}

type SolveResponse struct {
	Solved   bool      `json:"solved"`
	Solution [N][N]int `json:"solution"`
	Message  string    `json:"message"`
}

// --- Handlers ---

func solveHandler(w http.ResponseWriter, r *http.Request) {
	// Enable CORS for browser requests
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		return
	}

	var req SolveRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	board := req.Board
	success := solveSudoku(&board)

	resp := SolveResponse{
		Solved:   success,
		Solution: board,
	}
	if !success {
		resp.Message = "This puzzle cannot be solved."
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	// Route to serve the HTML file
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	// Route for the solve API
	http.HandleFunc("/solve", solveHandler)

	fmt.Println("Server active at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}