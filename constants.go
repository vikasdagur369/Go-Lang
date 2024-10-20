package main

import (
	"fmt"
)

// we can declare variable outside the fun main

const b = 3

//but a := 4 vala varible func main ke bahar declare nhi kar skte
func main() {
	//we can use const keyword for variables, which will be similar to js const

	const a = 4
	const name = "vikas"

	fmt.Println(name)
}
