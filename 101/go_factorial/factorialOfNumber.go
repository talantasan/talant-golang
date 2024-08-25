package main

type fact int

func calcFunc(f fact) (fact){
	result := fact(1)
	for i := range f{
		result *= (i+1)
	}
	return result
}