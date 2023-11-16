package main

import "fmt"

func main() {
	//path 1
	myVar := 10

	ptr1 := &myVar

	ptr2 := &ptr1

	fmt.Println(myVar, *ptr1, **ptr2)
	fmt.Println(&myVar, &ptr1, &ptr2)

	**ptr2--

	fmt.Println(myVar, *ptr1, **ptr2)

	fmt.Println()
	//path 2
	myArr := [...]int{10, 20, 30}

	fmt.Println(myArr)

	ptrArr := &myArr

	ptrArr[1] = 40

	fmt.Println(myArr)
}
