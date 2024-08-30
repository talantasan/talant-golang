package main

import (
	"fmt"
	"sort"
)

func main() {

	ages := []int{40, 42, 16, 34, 33, 50, 21}
	names := [5]string{"Talant", "Samara", "Meeriem", "Azhar", "Jasmine"}

	sort.Ints(ages) // Sorts in place
	fmt.Println(ages)

	index := sort.SearchInts(ages, 34) // Sort package, SearchInts method will return index position out of sorted sliece
	fmt.Println(index)


	sort.Strings(names[:])
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names[:], "Meeriem"))
}
