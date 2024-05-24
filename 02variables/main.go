package main

import "fmt"

func main() {
	var username string = "Alok"
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type: %T \n", isLoggedIn)

	var smallval uint8 = 252
	fmt.Println(smallval)
	fmt.Printf("variable is of type: %T \n", smallval)

	var smallflote float64 = 255.45536563
	fmt.Println(smallflote)
	fmt.Printf("variable is of type: %T \n", smallflote)

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("variable is of type: %T \n", anotherVariable)

	var website = "alok"
	fmt.Println(website)
	fmt.Printf("variable is of type: %T \n", website)

	numberOfUsers := 30000.0
	fmt.Println(numberOfUsers)
}
