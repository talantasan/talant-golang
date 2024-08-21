package main

import "fmt"

type books []string

func (b books) print(){
	for i, book := range b{
		fmt.Println(i, book)
	}
}

