package main

import "fmt"

func main() {
	var fruits []string
	fruits = []string{"banana", "coconut", "apple", "orange"}

	for i := 0; i < len(fruits); i++ {
		fmt.Println(fruits[i])
	}

	fmt.Println()

	matrix := [][]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}
}
