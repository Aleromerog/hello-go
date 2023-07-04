package main

import (
	"fmt"

	epicpizza "github.com/Aleromerog/michipizza"
)

func main() {
	// Get a greeting message and print it.

	message, err := epicpizza.RandomPizzaRecipe()

	if err != nil {
		fmt.Printf("%v", err)
	}

	fmt.Println("Ingredients: ")
	for _, elem := range message["ingredients"] {
		fmt.Println(elem)
	}

	fmt.Println("\nInstructions:")
	for i, elem := range message["instructions"] {
		fmt.Printf("%d. %v\n\n", i+1, elem)
	}
}
