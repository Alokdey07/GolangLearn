package main

import "fmt"

func main() {
	fmt.Println("welcome to array in go lang")

	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Apple"

	fmt.Println(fruitList)
	fmt.Println(len(fruitList))

	var vegList = [3]string{"Potato", "Beans", "Tomato"}
	fmt.Println("vegy list is: ", len(vegList), vegList)
}
