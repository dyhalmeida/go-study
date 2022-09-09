package main

import (
	"fmt"
)

func fibonnaci(position uint) uint {
	if (position <= 1) {
		return position
	}
	return fibonnaci(position - 2) + fibonnaci(position - 1)
}

func main(){
	fmt.Println(fibonnaci(8))

	position := uint(12)
	for i := uint(0); i < position; i++ {
		fmt.Print(fibonnaci(i), " ")
	}
}