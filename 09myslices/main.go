package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to video on slices")

	var fruitList = []string{"Apple", "Mango", "tomato", "Peach"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList)
	fmt.Println(fruitList)
	fruitList = append(fruitList, "Pomegranet", "Banana")

	fmt.Println(fruitList)

	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	highScore := make([]int, 4)

	highScore[0] = 234
	highScore[1] = 345
	highScore[2] = 465
	highScore[3] = 867

	highScore = append(highScore, 555, 666, 321)
	fmt.Println(highScore)

	sort.Ints(highScore)

	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted(highScore))

	//how to remove a value from slices baced on index
	var course = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(course)

	var index int = 2
	course = append(course[:index], course[index+1:]...)
	fmt.Println(course)
}
