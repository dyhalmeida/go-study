package main

import "fmt"

func main() {
	pointers()
	pointersWithStruct()
}

func pointers() {
	address := "Rua das Palmas, 07"
	fmt.Println(address)

	/* Ponteiros guardam a referência de memória da variável */
	diegoAddress := &address // aqui diegoAddress vira um ponteiro com a referência para o endereço de memória de address
	mayaraAddress := &address
	fmt.Println(diegoAddress, mayaraAddress) // aqui é exibido no terminal o endereço de memória
	fmt.Println(*diegoAddress, *mayaraAddress) // aqui é exibido o valor armazenado do endereço de memória

	address = "Caminho 26, 07"
	fmt.Println(address, diegoAddress, mayaraAddress)
	fmt.Println(address, *diegoAddress, *mayaraAddress)

	/*O valor padrão de um ponteiro sem atribuição é <Nil>*/
	var test *string
	fmt.Println(test) // <Nil>

}

func pointersWithStruct() {

	type Address struct {
		number int8
		name string
	}

	NewAddress := func(number int8, name string) (*Address) {
		return &Address{number: number, name: name}
	}

	// GetNumber := func (a Address) int8 {
	// 	return a.number
	// }

	// GetAddressName := func (a Address) string {
	// 	return a.name
	// }

	address := NewAddress(01, "Rua Vital Brasil")
	// fmt.Println(GetNumber(*address))
	// fmt.Println(GetAddressName(*address))

	marcosAddress := address
	ivoneAddress := address
	fmt.Println(marcosAddress, ivoneAddress)
	fmt.Println("Marcos mora: " + marcosAddress.name)
	fmt.Println("Ivone mora: " + ivoneAddress.name)

	address.name = "Rua Antônio Luis Garcez"
	fmt.Println("Marcos mora: " + marcosAddress.name)
	fmt.Println("Ivone mora: " + ivoneAddress.name)
}

func test() {
	age := 30
	fmt.Println(&age) // 0xc0000160d0

	var name *string
	fmt.Println(name) // <Nil>
	fmt.Println(&name)// 0xc00000e030
}
