package main

import "fmt"

func displayResult(num int, sq int, cb int) {
	fmt.Printf("Square of %v is %v \n", num, sq)
	fmt.Printf("Cube of %v is %v \n", num, cb)
	fmt.Println("----------------------------")
}

func calcResult(num int) (int, int) {
	return num * num, num * num *num 
}

func iterateNums(n []int, f func(int, int, int)) {
	for _, num := range n {
		sq, cb := calcResult(num)
		f(num, sq, cb)
	}
}

func main() {
	nums := []int{4, 5, 2, 100}
	iterateNums(nums, displayResult)
}