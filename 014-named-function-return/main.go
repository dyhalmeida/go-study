package main

import "fmt"

func main() {
	sumResult, subResult := sum(4, 8)
	fmt.Println(sumResult, subResult)
}

/* Funções com retorno nomeado, possui seu retorno declarado
e no escopo da função apenas a key return, já retorna o que foi especificado*/
func sum(n1, n2 int) (sumResult int, subResult int) {
	sumResult = n1 + n2
	subResult = n1 - n2
	return
}