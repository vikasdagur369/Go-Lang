package main

import "fmt"

func main() {

	// while loop
	var i = 1

	for i < 5 {
		fmt.Println(i)
		i++
	}

	// infinite loop
	//for{
	// write anything
	//}

	// traditional for loop

	for j := 1; j < 5; j++ {
		fmt.Println(j)
	}

	// isme i ki value se start hokr range-1 tak chalega
	for i := range 8 {
		fmt.Println(i)
	}
}
