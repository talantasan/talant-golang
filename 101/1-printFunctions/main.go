package main

import "fmt"

func main() {
	
	//PRint does not add a new line unlike Println
	fmt.Print("Hello world\n")

	// Formatting string
	var name string = "Talant"
	var age int = 35

	// format specifie _%
	// %v - value, %q - quote, %T - type etc.

	fmt.Printf("Hello all, my name is %v, I am %v years old \n", name, age)
	fmt.Printf("My name is %q and age is type of %T and float number is %0.3f \n", name, age, 23.4352345)
	
	//Sprintf -- save formatted string
	str := fmt.Sprintf("Hello all, my name is %v, I am %v years old", name, age)
	fmt.Println(str)
}