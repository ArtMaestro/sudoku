package main

import (
	"fmt"
)

type Board [9][9]int

func isSafe(board Board, row, col, num int) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == num || board[i][col] == num {
			return false
		}
	}

	startRow, startCol := row-row%3, col-col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i+startRow][j+startCol] == num {
				return false
			}
		}
	}

	return true
}

func solveSudoku(board Board) bool {
	empty := true
	var row, col int

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				row, col = i, j
				empty = false
				break
			}
		}
		if !empty {
			break
		}
	}

	if empty {
		return true
	}

	for num := 1; num <= 9; num++ {
		if isSafe(board, row, col, num) {
			board[row][col] = num
			if solveSudoku(board) {
				return true
			}
			board[row][col] = 0
		}
	}

	return false
}

func printBoard(board Board) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%d ", board[i][j])
		}
		fmt.Println()
	}
}

func main() {
	board := Board{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	fmt.Println("Исходная судоку:")
	printBoard(board)

	if solveSudoku(board) {
		fmt.Println("\n Решение:")
		printBoard(board)
	} else {
		fmt.Println("Упс, не выйдет")
	}
}

// func TestSolveSudoku(t *testing.T) {
// 	board := Board{
// 		{5, 3, 0, 0, 7, 0, 0, 0, 0},
// 		{6, 0, 0, 1, 9, 5, 0, 0, 0},
// 		{0, 9, 8, 0, 0, 0, 0, 6, 0},
// 		{8, 0, 0, 0, 6, 0, 0, 0, 3},
// 		{4, 0, 0, 8, 0, 3, 0, 0, 1},
// 		{7, 0, 0, 0, 2, 0, 0, 0, 6},
// 		{0, 6, 0, 0, 0, 0, 2, 8, 0},
// 		{0, 0, 0, 4, 1, 9, 0, 0, 5},
// 		{0, 0, 0, 0, 8, 0, 0, 7, 9},
// 	}

// 	expectedSolution := Board{
// 		{3, 1, 6, 5, 7, 8, 4, 9, 2},
// 		{5, 2, 9, 1, 3, 4, 7, 6, 8},
// 		{4, 8, 7, 6, 2, 9, 5, 3, 1},
// 		{2, 6, 3, 4, 1, 5, 9, 8, 7},
// 		{9, 7, 4, 8, 6, 3, 1, 2, 5},
// 		{8, 5, 1, 7, 9, 2, 6, 4, 3},
// 		{1, 3, 8, 9, 4, 7, 2, 5, 6},
// 		{6, 9, 2, 3, 5, 1, 8, 7, 4},
// 		{7, 4, 5, 2, 8, 6, 3, 1, 9},
// 	}

// 	solveSudoku(board)

// 	for i := 0; i < 9; i++ {
// 		for j := 0; j < 9; j++ {
// 			if board[i][j] != expectedSolution[i][j] {
// 				t.Errorf("Expected %d at (%d,%d) but got %d", expectedSolution[i][j], i, j, board[i][j])
// 			}
// 		}
// 	}
// }
