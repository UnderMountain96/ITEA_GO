package main

import "fmt"

func main() {
	var fruits [4]string
	fruits = [4]string{"banana", "coconut", "apple", "orange"}

	for i := 0; i < len(fruits); i++ {
		fmt.Println(fruits[i])
	}

	fmt.Println()

	matrix := [3][4]int{}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}
}
