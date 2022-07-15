package main

import "fmt"

func main(){

	/*Exemplo de utilização de estrutura condicional*/
	number := 12

	if number % 2 == 0 {
		fmt.Println("Par")
	} else {
		fmt.Println("Ímpar")
	}

	/*if com inicialização de variável*/
	if par := number; par > 10 {
		fmt.Println(par) // 12
		fmt.Println("É maior que dez")
	} else {
		fmt.Println("É menor que dez")
	}

	fmt.Println(checkAge())

}

func checkAge() string {

	age := 18

	if age >= 18 {
		return "É maior de idade"
	}

	return "É menor de idade"

}