package main

import "fmt"

func main() {

	// defining string variables
	var name string = "Talant"
	var country = "Kyrgyzstan"
	lastName := "Karmyshakov"

	fmt.Println(name+" "+lastName+" from "+country)

	// reassign values of variables
	name = "Samara"
	lastName = "Ataeva"
	country = "KG"
	fmt.Println(name+" "+lastName+" from "+country)

	// defining numbers
	var num_int1 int = 1
	var num_float float32 = 2.5

	fmt.Println(num_int1, num_float)

	// bits and memory
	var num1 int8 = 12 //in a specific range -128 to 128
	var num2 int32 = 234234324 // if exceed the range you will see error underlined
	var num3 float32 = 234234234234234243.324234234234234234234234234234234234234234234234234234234234

	fmt.Println(num1, num2, num3)
}