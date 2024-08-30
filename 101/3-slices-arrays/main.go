package main

import "fmt"

func main() {
	
	// Arrays - Array with declared length, array must have fixed length unlike slices
	var nums1 [3]int = [3]int{20, 30, 40} 
	var nums2 = [3]int{100, 200, 500}
	names := [5]string{"Talant", "Samara", "Meeriem", "Azhar", "Jasmine"}

	for _, num := range nums1{
		fmt.Println(num)
	}

	for _, num := range nums2{
		fmt.Println(num)
	}

	i := 0
	for i < len(names){
		fmt.Println(names[i])
		i++
	}
	fmt.Println("-----------END OF ARRAYS-------------")
	
	// 2. Slices, length is not defined
	
	var scores = []int{100, 40, 60}

	for i := 0; i < len(scores); i++ {
		fmt.Println(scores[i])
	}

	// append value to slice
	scores = append(scores, 1000)

	for i := 0; i < len(scores); i++ {
		fmt.Println(scores[i])
	}


	// slice ranges
	fmt.Println("-----------SLICE Ranges-------------")
	newRange := scores[2:] // from index 2 to last index

		for i := 0; i < len(newRange); i++ {
		fmt.Println(newRange[i])
	}
	
}