package main

import "fmt"

func main() {

	func () {
		fmt.Println("Hello, i am a anonymou function")
	}()


	func (name string) {
		fmt.Sprintf("Hello, %s", name)
	}("Diego Almeida")

	birthdayYear := func (birthdayYear, currentYear int) int {
		return currentYear - birthdayYear
	}(1992, 2022)

	fmt.Println(birthdayYear)

}