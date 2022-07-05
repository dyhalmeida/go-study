package main

import "fmt"

type Address struct {
	city string
	state string
}

type User struct {
	name string
	age int
	address Address
}

func main() {
	/* O structs possui basicamente 3 formas de ser inicializados*/
	struct1()
	struct2()
	struct3()
}

func struct1() {
	/* Inicializado por propriedades individuais*/
	var user User
	user.name = "Diego"
	user.age = 30

	var address Address
	address.city = "Mata de São João"
	address.state = "BA"

	user.address = address

	fmt.Println(user)
}

func struct2() {

	/* Inicializado por inferência implicita de tipo */
	address := Address{"Mata de São João", "BA"}
	user := User{"Laura", 6, address}
	fmt.Println(user)

}

func struct3() {
	/* Inicializado por inferência implicita e omitindo valores*/
	/* omitindo o address */
	user := User{name: "Mayara", age: 26}
	fmt.Println(user)

}