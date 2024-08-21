package main

import "fmt"

type common_slice []string

func (cs common_slice) print(){
	for i, val := range cs{
		fmt.Println(i, val)
	}
}

