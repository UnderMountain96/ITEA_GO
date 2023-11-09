package task

import "fmt"

func Variables() {
	var boolVar bool

	fmt.Println(boolVar)

	boolVar = true

	fmt.Println(boolVar)

	var strVar string = "some text"

	fmt.Println(strVar)
}
