package main

import (
	"fmt"
	"math"
)

func rotate(matrix [][]int) {
	n := len(matrix[0])
	round := int(math.Ceil(float64(n) / 2))

	for r := 0; r < round; r++ {
		length := n - (r * 2)

		if length == 1 {
			break
		}

		startPos := r
		endPos := r + length - 1
		temp := 0

		for i := startPos; i < endPos; i++ {
			//top to left
			nextCell := matrix[i][endPos]
			matrix[i][endPos] = matrix[startPos][i]

			//left to bottom
			nextCell, matrix[endPos][endPos-temp] = matrix[endPos][endPos-temp], nextCell

			//bottom to right
			nextCell, matrix[endPos-temp][startPos] = matrix[endPos-temp][startPos], nextCell

			//right to top
			matrix[startPos][i] = nextCell

			// nextCell := matrix[startPos][i]

			// //right to top
			// matrix[startPos][i] = matrix[endPos-temp][startPos]

			// //bottom to right
			// matrix[endPos-temp][startPos] = matrix[endPos][endPos-temp]

			// //left to bottom
			// matrix[endPos][endPos-temp] = matrix[i][endPos]

			// //top to left
			// matrix[i][endPos] = matrix[startPos][i]

			// temp++
		}
	}

	fmt.Println(matrix)

}

// {{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
func main() {
	data := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	// data := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(data)
	rotate(data)
}
