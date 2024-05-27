package main

import (
	"fmt"
)

func main() {
	fmt.Println("maps in golang")

	language := make(map[string]string)

	language["JS"] = "Javascript"
	language["RB"] = "Ruby"
	language["PY"] = "Python"

	fmt.Println("List of all languages: ", language)
	fmt.Println("Js value: ", language["JS"])

	delete(language, "RB")

	fmt.Println("List of all language: ", language)

	// loops are interesting in golang
	for key, value := range language {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

	for _, value := range language {
		fmt.Printf("For key v, value is %v\n", value)
	}

}
