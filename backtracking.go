package main

import (
	"fmt"
	"os"
)

const boardSize = 9

func isValid(board [][]int, row, col, num int) bool {
	for i := 0; i < boardSize; i++ {
		if board[row][i] == num || board[i][col] == num {
			return false
		}
	}

	startRow, startCol := row-row%3, col-col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[startRow+i][startCol+j] == num {
				return false
			}
		}
	}

	return true
}

func solveSudoku(board [][]int) bool {
	for row := 0; row < boardSize; row++ {
		for col := 0; col < boardSize; col++ {
			if board[row][col] == 0 {
				for num := 1; num <= boardSize; num++ {
					if isValid(board, row, col, num) {
						board[row][col] = num
						if solveSudoku(board) {
							return true
						}
						board[row][col] = 0
					}
				}
				return false
			}
		}
	}
	return true
}

func printBoard(board [][]int) {
	for _, row := range board {
		for _, num := range row {
			fmt.Printf("%d ", num)
		}
		fmt.Println()
	}
}

func main() {
	args := os.Args[1:]
	if len(args) != boardSize {
		fmt.Println("Error")
		return
	}

	board := make([][]int, boardSize)
	for i := range board {
		board[i] = make([]int, boardSize)
		for j, char := range args[i] {
			if char == '.' {
				board[i][j] = 0
			} else {
				board[i][j] = int(char - '0')
			}
		}
	}

	if solveSudoku(board) {
		printBoard(board)
	} else {
		fmt.Println("Error")
	}
}
