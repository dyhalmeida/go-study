package main

import "fmt"

/*
	A cláusula Defer, adia a execução de uma função.
	Se uma função estiver sendo chamada com a cláusula Defer, dentro
	de uma função se retorno. Ela será a última a ser executada.
	Se ela estiver sendo chamada dentro de uma função com retorno, ela será
	executada sempre antes de um return.
*/
func main() {

	defer start()
	stop()

	avg(7, 8)

}

func start() {
	fmt.Println("Start code")
}

func stop() {
	fmt.Println("Stop code")
}

func avg(n1, n2 float32) bool {

	defer fmt.Println("Resultado da média")

	average := n1 + n2 / 2

	if (average >= 6) {
		fmt.Println("Condição V")
		return true
	}
	fmt.Println("Condição F")
	return false
	
}