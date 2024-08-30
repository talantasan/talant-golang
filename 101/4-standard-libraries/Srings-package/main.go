package main

import (
	"fmt"
	"strings"
)

func main() {
	greeting := "Hello friends! Welcome to Talant's workspace"

	//strings library, Contains method
	fmt.Println(strings.Contains(greeting, "Talant")) // Retruns boolean value True | False
	fmt.Println(strings.Contains(greeting, "Aibek"))

	//strings library, Replace method
	fmt.Println(strings.ReplaceAll(greeting, "Talant", "Aibek"))

	//strings library, ToUpper method
	fmt.Println(strings.ToUpper(greeting))

	//strings library, Index method
	fmt.Println(strings.Index(greeting, "COME"))

	//strings library, Split method
	fmt.Println(strings.Split(greeting, " "))
}