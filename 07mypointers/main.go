package main

import "fmt"

func main() {
	fmt.Println("hello")

	//var one int = 2
	// var abc *int

	// fmt.Println("Value of abc: ", abc)

	myNumber := 23
	var ptr = &myNumber
	var ab = *ptr

	fmt.Println("Value of ptr: ", ptr)
	fmt.Println("Value of ptr: ", *ptr)
	fmt.Println("Value of ab: ", ab)
	fmt.Println(ab)

	*ptr = *ptr * 2
	fmt.Println("new Value of ptr: ", ptr)
	fmt.Println("new Value of ab: ", ab)
	fmt.Println("new Value of mynum: ", myNumber)
}
