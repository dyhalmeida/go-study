package main

import "fmt"

func main() {
	arithmeticOperators()
	fmt.Println("========================================")
	relationalOperators()
	fmt.Println("========================================")
	assignmentOperators()
	fmt.Println("========================================")
	logicalOperators()
	fmt.Println("========================================")
	unaryOperators()
}


func arithmeticOperators() {
	/*
	Os operadores aritméticos são
	*/
	fmt.Println(10 + 10)
	fmt.Println(10 - 10)
	fmt.Println(10 * 10)
	fmt.Println(10 / 10)
	fmt.Println(10 % 10)

	/* Observação: Não se poder fazer operações com tipos de dados diferentes como por exemplo
	int8 + int16*/

}

func relationalOperators() {
	fmt.Println(10 > 5)
	fmt.Println(10 < 5)
	fmt.Println(10 >= 5)
	fmt.Println(10 <= 5)
	fmt.Println(10 != 5)
	fmt.Println(10 == 5)
}

func assignmentOperators() {
	/* Atribuição sem inferência de tipo*/
	var name string = "Diego"
	fmt.Println(name)

	/* Atribuição com inferência de tipo*/
	lastname := "Almeida"
	fmt.Println(lastname)

}

func logicalOperators() {
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Println(!false)
}

func unaryOperators() {
	var counter int = 0
	counter++
	fmt.Println(counter)
	counter--
	fmt.Println(counter)
	counter += 10
	fmt.Println(counter)
	counter -= 10
	fmt.Println(counter)
	counter *= 10
	fmt.Println(counter)
}