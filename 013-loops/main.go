package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		No Golang só existe o FOR para executar loops
	*/
	loopA()
	loopB()
	loopC()
	loopD()
	loopE()
	loopF()
}


func loopA() {
	i := 0

	for i < 10 {
		i++
		fmt.Println("Simulado Do While: ", i)
		time.Sleep(time.Second)
	}
}

func loopB() {

	/* For com declaração e inicialização de varável */

	for i := 0; i <= 10; i++ {
		fmt.Println("Loop: ", i)
		time.Sleep(time.Second)
	}
}

func loopC() {

	/* Iterando slice e array*/

	// Slice
	names := []string{"Diego", "Laura", "Mayara"}
	for index, name := range names {
		fmt.Println(index, name)
		time.Sleep(time.Second)
	}

	// Array e ignorando primeiro valor retornado de range
	ages := [3]int{30, 26, 7}
	for _, age := range ages {
		fmt.Println(age)
		time.Sleep(time.Second)
	}
}

func loopD() {
	/* Iterando map*/
	people := map[string]string {
		"name": "Diego",
		"lastname": "Almeida",
	}
	for key, value := range people {
		fmt.Println(key, value)
		time.Sleep(time.Second)
	}
}

func loopE() {

	/* Iterando string
	Obs: Numa iteração de string o valor retornado é o código da letra na tabela ASCI
	*/

	word := "A long long ago..."
	for _, letter := range word {
		fmt.Println(letter, string(letter))
		time.Sleep(time.Second)
	}
}

func loopF() {
	/* Loop infinito */
	for {
		fmt.Println("Loop infinito")
		time.Sleep(time.Second)
	}
}