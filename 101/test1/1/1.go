package main

import "fmt"

func main() {
	
	// String variable
	var name string = "Talant"
	var lastName = "Karmyshakov"
	country := "KG"

	// Numeric variable, int || float
	var num1 int = 12341234
	var num2 = 234234
	num3 := 23.13265

	fmt.Println(name,lastName, country )

	fmt.Printf("My name is %q last name is %v \n", name, lastName)
	fmt.Printf("Integer number %v and %v, float is %v \n", num1, num2, num3)

	
}