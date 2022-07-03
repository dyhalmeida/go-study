package main

import "fmt"

func main() {
	resultSum := sum(10, 20)
	fmt.Println(resultSum)

	sayName()

	sum, div := calculations(40, 20)
	fmt.Println(sum, div)

	/* Para ignora um dos retornos da função pode-se utilizar o underline(_)*/
	sum2, _ := calculations(100, 80)
	fmt.Println(sum2)

	_, div2 := calculations(100, 80)
	fmt.Println(div2)
}

/* Funções em Go com parâmetros e retorno */
func sum(x int, y int) int {
	return x + y
}

/* Funções como variáveis */
var sayName = func() {
	fmt.Println("Meu nome é Diego Almeida")
}

/* Funções com multiplos retornos*/
var calculations = func(x float64, y float64) (float64, float64) {
	sum := x + y
	div := x / y
	return sum, div
}