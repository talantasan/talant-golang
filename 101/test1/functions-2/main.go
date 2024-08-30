package main

import "fmt"

func calcFact(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func iterateFacts(s []int, f func(int, int)) {
	for _, val := range s {
		r := calcFact(val)
		f(val, r)
	}
}

func displayMessage(n1 int, r2 int) {
	fmt.Printf("Factorial of %v is %v \n", n1, r2)
}

func main() {
	nums := []int{3, 2, 5, 4, 7}
	iterateFacts(nums, displayMessage)
}

