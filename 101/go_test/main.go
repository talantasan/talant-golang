package main

import "fmt"

func main(){
	// return message
	fmt.Println("Hello World")
	
	// call function
	fmt.Println(newMessage())

	//loop 
	series := []string{"House of Cards", "Lord of Rings", "Dexter"}
	series = append(series, "Doctor Who")
	for i, serie := range series{
		fmt.Println(i, serie)
	}

	fmt.Println("--------------------")

	// 
	books := books{"Kelkel", "Broken Sword", "Jamila"}
	books.print()
}

func newMessage() string{
	return "New message from function"
}
