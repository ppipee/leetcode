package main

import (
	"fmt"
)

// func getCellIndex(i int, j int) {
// 	return (i * 9) + j
// }

// func findAllCarry(board [][]byte) {
// 	carryBoard := [][][]string{}

// 	for i := 0; i < 9; i++ {
// 		for j := 0; j < 9; j++ {

// 		}
// 	}
// }

func createEmptyBoard() {
	board := [][]string{}

	for i := 0; i < 9; i++ {
		row := []string{}

		for j := 0; j < 9; j++ {
			row = append(row, string(j))
		}

		board = append(board, row)
	}
	fmt.Println(board)

	return board
}

func solveSudoku(board [][]byte) {
	isSolved := false
	// carryBoard := findAllCarry(board)
	fmt.Println(createEmptyBoard())
}

func main() {
	board := [][]byte{{"5", "3", ".", ".", "7", ".", ".", ".", "."}, {"6", ".", ".", "1", "9", "5", ".", ".", "."}, {".", "9", "8", ".", ".", ".", ".", "6", "."}, {"8", ".", ".", ".", "6", ".", ".", ".", "3"}, {"4", ".", ".", "8", ".", "3", ".", ".", "1"}, {"7", ".", ".", ".", "2", ".", ".", ".", "6"}, {".", "6", ".", ".", ".", ".", "2", "8", "."}, {".", ".", ".", "4", "1", "9", ".", ".", "5"}, {".", ".", ".", ".", "8", ".", ".", "7", "9"}}

	fmt.Println(solveSudoku(board))
}
