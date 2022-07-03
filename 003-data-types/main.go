package main

import (
	"errors"
	"fmt"
)

func main() {
	integers()
	floats()
	strings()
	booleans()
	singleQuotes()
	errorDataType()
}

func integers() {
	/*
	Os tipos de dados inteiros em Golang são:
	[int8, int16, int32, int64]
	A diferença entre eles é a quantidade de bits e o tamanho do número
	Já o tipo int  apenas, é definido sua quantidade de bit de acordo com a arquitetura do computador
	*/
	var age int8 = 30
	fmt.Println("Minha idade:", age)

	/*O valor padrão de uma variável int sem atribuição é 0(zero)*/
}

func floats() {
	/*
	Os tipos de dados decimais em Golang são:
	[float32, float64]
	A diferença entre eles é a quantidade de bits e o tamanho do número
	Já o tipo float apenas, é definido apenas por inferência de tipo ou tipagem implicita e 
	sua quantidade de bit é de acordo com a arquitetura do computador
	*/
	var height float32 = 1.75
	fmt.Println("Minha altura:", height)

	/*O valor padrão de uma variável float sem atribuição é 0(zero)*/

}

func strings() {
	/*
	Os tipos de dados para textos em Golang é:
	[string]. Uma string em Go deve ser entre aspas duplas
	*/
	var name string = "Diego Almeida"
	fmt.Println("Meu nome:", name)

	/*O valor padrão de uma variável string sem atribuição é "" string vazia*/
}

func booleans() {
	/*
	Os tipos de dados para valores lógicos em Golang é:
	[bool].
	*/

	var isMarried bool = true
	fmt.Println("Sou casado?", isMarried)

	/*O valor padrão de uma variável bool sem atribuição é false*/

}

func errorDataType() {

	/*
	O tipo de dados para erros em Golang é:
	[error]. Para atribuir ou criar um error, deve utilizar errors.New()
	*/

	var error error = errors.New("Internal server error")
	fmt.Println(error)

	/*O valor padrão de uma variável error sem atribuição é <Nil>*/

}

func singleQuotes() {
	/*
	Uma variável declarada que recebe um valor entres aspas simples com inferência de tipo,
	é do tipo int e seu valor recebe a refência da tabela ASCII do conteúdo entre aspas simples
	atribuido.
	*/
	var char = 'A'
	fmt.Println(char)
}