package main

import (
	"fmt"
	"math"
)

func sayGreeting(n string) {
	fmt.Printf("Good morning %v \n", n)
}

func calcSum(n1 int, n2 int) (int, int){
	return n1+n2, n1*n2
}

func iterateNames(n []string, f func(string)) {
	for _, val := range n{
		f(val)
	}
}

func circleArea(r float64) (float64){
	return math.Pi*r*r
}

func main() {

	n1, n2 := 2, 5
	sayGreeting("Talant")
	s, p := calcSum(n1, n2)
	
	fmt.Printf("%v+%v=%v \n", n1, n2, s)
	fmt.Printf("%v*%v=%v \n", n1, n2, p)

	names := []string{"Samara", "Meeriem", "Azhar", "Jasmine"}
	
	iterateNames(names, sayGreeting)

	a1 := circleArea(10.4)
	fmt.Println(a1)

}


