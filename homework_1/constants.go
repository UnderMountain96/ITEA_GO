package main

import "fmt"

const intConst = 2 * 2

type WeekDay int

const (
	Monday WeekDay = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func main() {
	fmt.Println(intConst)
}
