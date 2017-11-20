package main

import "fmt"

func main() {
	fmt.Println(rotate())
}

func rotate() [][]int {
	var matrix = [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}

	crossDown(reverse(matrix))

	return matrix
}

func reverse(matrix [][]int) [][]int {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		note := matrix[n-i-1]
		matrix[n-i-1] = matrix[i]
		matrix[i] = note
	}

	return matrix
}

func crossDown(matrix [][]int) [][]int {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i <= j {
				continue
			}
			tmp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = tmp
		}
	}

	return matrix
}
