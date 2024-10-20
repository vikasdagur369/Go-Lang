package main

import (
	"fmt"
)

func main() {
	//variables declaration types
	// golang me variables ko use karna jaruri hai

	// declaring using var keyword and type
	var b string = "hello"
	var a int = 7
	fmt.Println(a, b)

	//declaring using only var keyword

	var name = "vikas"
	var surName = "dagur"

	fmt.Println(name + surName)

	//direct declare

	name1 := "shikha"
	name2 := "dagur"

	fmt.Println(name1 + name2)

}
