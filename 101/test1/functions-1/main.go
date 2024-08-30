package main

import (
	"fmt"
	"math"
)

func sayGreeting(s1 string) {
	fmt.Println("Hello " + s1 + " welcome")
}

func loopUsers(names []string, f func(string)) {
	for _, val := range names {
		f(val)
	}
}

func cirArea(r float64) float64 {
	return math.Pi * r * r
}

func main() {
	// sayGreeting("Talant")
	names := []string{"Talant", "Samara", "Meeriem", "Azhar", "Jasmine"}

	loopUsers(names, sayGreeting)

	rad := 3.4
	fmt.Printf("Area of the cirle is %0.4f", cirArea(rad))
	fmt.Println()
}
