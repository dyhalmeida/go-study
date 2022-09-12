package main

import "fmt"

func closure() func() {
	text := "Hello closure"

	fn := func() {
		fmt.Println(text)
	}

	return fn
}

func main() {
	fmt.Println("Hello main")
	fn := closure()
	fn()
}