package main

import "fmt"

func main(){
//1.--------------
	// x := 1
	// for x<=5{
	// 	fmt.Println(x)
	// 	x++
	// }
//2.--------------

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

//3.--------------
	
	// names := []string{"Talant", "Samara", "Meeriem", "Azhar", "Jasmine"}

	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }

//4.--------------
	names := []string{"Talant", "Samara", "Meeriem", "Azhar", "Jasmine"}
	for index, value := range names {
		
		fmt.Printf("index is %v, value is %v \n", index, value)
	
	}

}