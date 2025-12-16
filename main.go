package main

import (
	"fmt"
)

// Global constants for the Sudoku grid size
const N = 9

// isSafe checks if it's safe to place 'num' at board[row][col]
// It checks the row, column, and the 3x3 subgrid.
func isSafe(board *[N][N]int, row, col, num int) bool {
	// 1. Check Row
	for c := 0; c < N; c++ {
		if board[row][c] == num {
			return false
		}
	}

	// 2. Check Column
	for r := 0; r < N; r++ {
		if board[r][col] == num {
			return false
		}
	}

	// 3. Check 3x3 Box
	startRow := row - row%3
	startCol := col - col%3

	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			if board[startRow+r][startCol+c] == num {
				return false
			}
		}
	}

	// If num is safe in all three checks, return true
	return true
}

// solveSudoku implements the backtracking algorithm.
// It returns true if a solution is found, false otherwise.
func solveSudoku(board *[N][N]int) bool {
	// Find the next empty cell (represented by 0)
	var row, col int
	foundEmpty := false

	for r := 0; r < N; r++ {
		for c := 0; c < N; c++ {
			if board[r][c] == 0 {
				row, col = r, c
				foundEmpty = true
				break
			}
		}
		if foundEmpty {
			break
		}
	}

	// Base Case: If no empty cell was found, the board is solved.
	if !foundEmpty {
		return true
	}

	// Try numbers 1 through 9
	for num := 1; num <= 9; num++ {
		if isSafe(board, row, col, num) {
			// Place the number
			board[row][col] = num

			// Recurse: Try to solve the rest of the puzzle
			if solveSudoku(board) {
				return true // Solution found! Propagate success up the stack
			}

			// Backtrack: If the number didn't lead to a solution, reset the cell to 0
			board[row][col] = 0
		}
	}

	// No number from 1-9 worked for this cell. Trigger backtrack on the previous cell.
	return false
}

// printBoard prints the Sudoku board in a readable format.
func printBoard(board *[N][N]int) {
	for r := 0; r < N; r++ {
		for c := 0; c < N; c++ {
			fmt.Printf("%d ", board[r][c])
			if (c+1)%3 == 0 && c != N-1 {
				fmt.Print("| ")
			}
		}
		fmt.Println()
		if (r+1)%3 == 0 && r != N-1 {
			fmt.Println("---------------------")
		}
	}
}

func main() {
	// Sample Sudoku puzzle (0 represents an empty cell)
	board := [N][N]int{
		{2, 0, 7, 0, 0, 4, 0, 8, 0}, 
		{0, 0, 0, 5, 0, 3, 9, 0, 0}, 
		{0, 3, 0, 6, 0, 0, 0, 2, 7}, 
		{0, 0, 0, 0, 0, 1, 4, 0, 0}, 
		{7, 5, 0, 0, 0, 0, 0, 1, 9}, 
		{0, 0, 9, 8, 0, 0, 0, 0, 0}, 
		{4, 2, 0, 0, 0, 6, 0, 9, 0}, 
		{0, 0, 6, 3, 0, 5, 0, 0, 0}, 
		{0, 9, 0, 1, 0, 0, 7, 0, 2},

	}

	fmt.Println("--- Unsolved Sudoku Board ---")
	printBoard(&board)
	fmt.Println()

	if solveSudoku(&board) {
		fmt.Println("--- Solved Sudoku Board ---")
		printBoard(&board)
	} else {
		fmt.Println("No solution exists for the given Sudoku puzzle.")
	}
}