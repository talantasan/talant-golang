package main

import "fmt"

type squareCube int

func calcSqCube(sq squareCube, x squareCube) (squareCube){
	result := squareCube(1)

	for  := range x{
		result *= sq
	}
	return result
}

func (sq squareCube) print(){
	fmt.Println(sq)
}

