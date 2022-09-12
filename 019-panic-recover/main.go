package main

import "fmt"

func recovery() {
	if r := recover(); r != nil {
		fmt.Println("Programa recuperado")
	}
}

func avg(n1, n2 float64) bool {
	defer recovery()
	average := (n1 + n2) / 2

	if average > 6 {
		return true
	} else if average < 6 {
		return false
	}

	/*
		A função panic gera um pánico no programa e mata o processo em execução.
		Antes do panic ser executado, todas as funções defer são executadas, sendo
		assim, podemos usar uma função defer com o recover() para tentar recuperar
		uma execução que entrou em pânico
	*/
	panic("A média é exatamente 6")
}

func main() {
	fmt.Println(avg(6, 7))
	fmt.Println("Programa excutado")

	fmt.Println(avg(6, 6))
	fmt.Println("Programa excutado")
}