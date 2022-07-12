package main

import "fmt"

func main() {
	arrays()
	slice()
}


func arrays() {
	/* Formas de se declarar e inicializar um array*/

	var numbersOfMegaSena [6]int
	/* O valor default de um array vazio do tipo int é zero(0) para cada posição*/
	fmt.Println(numbersOfMegaSena)

	numbersOfMegaSena[0] = 23
	numbersOfMegaSena[1] = 45
	numbersOfMegaSena[2] = 33
	numbersOfMegaSena[3] = 2
	numbersOfMegaSena[4] = 9
	numbersOfMegaSena[5] = 10
	fmt.Println(numbersOfMegaSena)

	/* O valor default de um array vazio do tipo string é uma string vazia para cada posição*/
	fruits := [...]string{"Banana", "Maçã", "Uva", "Goiaba"} // Tamano do array é definido pela quantidade de itens que recebe
	fmt.Println(fruits)


	months := [12]string{"Jan","Fev", "Mar", "Abr", "Maio", "Jun", "Jul", "Ago", "Set", "Out", "Nov", "Dez"}
	fmt.Println(months)
}

func slice() {
	/* Slice aponta para um array*/
	familyNames := []string{"Diego", "Laura", "May"}
	fmt.Println(familyNames)

	/* Para adicionar um novo item no slice, utiliza-se a função append, que retorna um novo slice
	com o item informado, a partir do slice anterior*/
	familyNames = append(familyNames, "Marcos")
	fmt.Println(familyNames)

	/*
		Um slice é um ponteiro para uma parte do array
	*/
	slice := familyNames[0:2]
	fmt.Println(slice)

	/*Slice é um ponteiro para uma parte do array, tanto é que se alterar o array, 
	o slice que está apontando pra essa fatia do array também sofre a alteração*/
	numbers := [4]int{1,2,3,4,}
	fmt.Println(numbers)
	sliceNumber := numbers[0:2]
	fmt.Println(sliceNumber)
	numbers[0] = 10
	fmt.Println(sliceNumber)
	fmt.Println("-----------------------------------------------------------------")

	/*
	Slice mais detalhado:
	Quando um slice é criado, o Go inicializa um espaço na memória como Array com a capacidade
	máxima igual o tamanho do Slice. Digamos que o slice tem 2 posições, o espaço na memória
	alocado, que o slice está apontando será um array de capacidade total igual a 2.
	Mas quando o slice estoura sua capacidade, ou seja, adicionando mais um item, o GO reserva outro
	espaço na memória para um novo Array com os 2 itens e sua capacidade total sobe para 4. Esse procedimento
	interno se repete sucessivamente sempre que o slice chegar na sua capacidade total.
	*/

	internalSlice := make([]int8, 2, 3)
	internalSlice[0] = 1
	internalSlice[1] = 2

	fmt.Println(len(internalSlice)) // tamanho = 2
	fmt.Println(cap(internalSlice)) // capacidade = 3
	fmt.Println(internalSlice) // output [1,2]

	internalSlice = append(internalSlice, 3) // Iguala acapacidade
	internalSlice = append(internalSlice, 4) // Passa da capacidade

	fmt.Println(len(internalSlice)) // tamanho = 4
	fmt.Println(cap(internalSlice)) // capacidade = 8
	fmt.Println(internalSlice) // output [1,2,3,4]

	fmt.Println("---------------------------------------------")
	internalSlice2 := make([]int8, 2)
	fmt.Println(internalSlice2) // [0,0]
	fmt.Println(len(internalSlice2)) // tamanho = 2
	fmt.Println(cap(internalSlice2)) // capacidade = 2
	internalSlice2 = append(internalSlice2, 1) // Passa da capacidade
	
	fmt.Println(internalSlice2) // [0,0,1]
	fmt.Println(len(internalSlice2)) // tamanho = 3
	fmt.Println(cap(internalSlice2)) // capacidade = 6
	

}