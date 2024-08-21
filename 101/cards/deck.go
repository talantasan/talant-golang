package main

import "fmt"

// create a new type of deck
// which is slice of strings

type common_slice []string

func (cs common_slice) print() {
	for i, val := range cs {
		fmt.Println(i, val)
	}
}

