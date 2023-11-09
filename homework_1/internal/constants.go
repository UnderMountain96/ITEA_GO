package task

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

func Constants() {
	fmt.Println(intConst)
}
