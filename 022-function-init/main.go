package main

import "fmt"

var n int

/*
	A função init sempre é executada antes da main, independente de onde ela esteja declarada no arquivo.
	Além disso, podemos várias funções init no projeto, uma por arquivo.
*/
func init() {
	fmt.Println("Executando init")
	n = 20
}

func main()  {
	fmt.Println("Executando main")
	fmt.Println(n)
}