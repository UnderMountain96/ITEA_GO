package main

import "fmt"

func main() {
	//path 1
	myVar := 10

	ptr1 := &myVar

	ptr2 := ptr1

	fmt.Println(myVar, *ptr1, *ptr2)

	*ptr2--

	fmt.Println(myVar, *ptr1, *ptr2)

}
