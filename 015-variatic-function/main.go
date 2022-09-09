package main

import "fmt"

func main() {
	total := sum(4, 4)
	fmt.Println(total)

	write("Esse é o número", 10, 20, 30, 40 ,50)
}

func sum(nums ...int) int {
	fmt.Println(nums)
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func write(text string, nums ...int) {
	for _, num := range nums {
		fmt.Println(text, num)
	}
}