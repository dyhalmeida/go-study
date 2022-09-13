package main

import "fmt"

func generic(t interface{}) {
	fmt.Println(t)
}

func main() {
	generic("Diego")
	generic(10)
	generic(true)
}