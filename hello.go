package main

import (
	"fmt"
)

func main() {
	// Get a greeting message and print it.
	message := Hello("Catalog")
	fmt.Println(message)
}

func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
