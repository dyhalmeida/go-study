package main

import "fmt"

func invertSignal(value int) int {
	return value * -1
}

/*
	Para acessar o valor do endereço de memória do ponteiro, utiliza-se o *
*/
func invertSignalWithPointers(value *int) {
	*value = *value * -1
}

func main() {

	n1 := 10
	n1Inverted := invertSignal(n1)
	fmt.Println(n1)
	fmt.Println(n1Inverted)

	n2 := 20
	fmt.Println(n2)

	/*
		Para acessar a referência de memória do valor primitivo, utiliza-se o &
	*/
	invertSignalWithPointers(&n2)
	fmt.Println(n2)

}