package main

import "fmt"

func main() {
	age := 29
	names := [5]string{"Talant", "Samara", "Meeriem", "Azhar", "Jasmine"}

	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 45)
	fmt.Println(age != 45)

	if age < 30 {
		fmt.Println("Age is less than 30")
	} else if age < 40 {
		fmt.Println("Age is less than 40")
	} else {
		fmt.Println("Age is not less than 40")
	}

	for index, value := range names{
		if index == 1 {
			fmt.Println("Continuing at pos ", index)
			continue
		}

		if index > 2 {
			fmt.Println("Ending here ...")
			break
		}

		fmt.Printf("the index at pos %v and value is %v \n", index, value)
	}

}