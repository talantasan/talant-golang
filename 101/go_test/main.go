package main

import "fmt"

func main(){
	// return message
	fmt.Println("Hello World")
	
	// call function
	fmt.Println(newMessage())

	//loop
	sum := 0
	for i := range 10{
		sum += i
		fmt.Println(i, sum)
	}

	//loop-slice 
	series := []string{"House of Cards", "Lord of Rings", "Dexter"}
	series = append(series, "Doctor Who")
	for i, serie := range series{
		fmt.Println(i, serie)
	}

	fmt.Println("--------------------")

	// 
	books := common_slice{"Kelkel", "Broken Sword", "Jamila"}
	books.print()

	fmt.Println("--------------------")

	students := common_slice{"Kapar", "Aibek", "Nurlan", "Junai"}
	students.print()

	
}

func newMessage() string{
	return "New message from function"
}
