package main

import "fmt"

func main() {
	/* Declaração explicita*/
	var name string = "Diego"
	var (
		lastname string = "Almeida"
		age int = 30
	)

	fmt.Println(name, lastname, age)

	/* Declaração implicita*/
	city := "Camaçari"
	address, num := "Rua das Palmas", 7
	cep, state := "42808139", "BA"

	fmt.Println(city, address, num)
	fmt.Println(cep, state)

	/* Constante segue a mesma regra de declaração de variáveis e como na maioria das linguagem não pode ter
	seu valor alterado depois de atribuido*/
	const yearOfBirth int = 1996
	fmt.Println(yearOfBirth)
}