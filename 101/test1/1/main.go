package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	var name string = "Talantbek Karmyshakov DevOps Engineer, Talantbek is very ambitious guy"
	var empYear = 4
	country := "Kyrgyzstan"

	fmt.Println(name)
	fmt.Println(empYear, country)

	newSlice := []string{}
	newSlice = append(newSlice, strings.Split(name, " ")...)
	fmt.Println(len(newSlice))

	//toLower
	fmt.Println(strings.ToLower(name))

	// Contains
	if strings.Contains(name, "SRE Engineer") {
		fmt.Println("We will consider to hire him")
	}

	// replaceAll
	fmt.Println(strings.ReplaceAll(name, "bek", "bai"))	
	fmt.Println(name)

	// Index
	fmt.Println(strings.LastIndex(name, "Talant"))

	fmt.Println("----------------------------------")

	// Sort Package

	ages := []int{40, 42, 16, 34, 33, 50, 21}
	names := [5]string{"Talant", "Samara", "Meeriem", "Azhar", "Jasmine"}

	
	index := sort.SearchInts(ages, 42)
	fmt.Println(ages)
	fmt.Println(index)

	sort.Ints(ages)
	sort.Strings(names[:])
	
	fmt.Println(names)
	fmt.Println(ages)



}