package main

import "fmt"

/*
	O Go não possui herança, porém a forma de herdar props de uma struct em outra struct
	é colocando a "struct pai" como propriedade da "struct filha", dessa forma a "struct filha",
	passará a ter também as props da "struct pai"
*/

type People struct {
	name string
	lastname string
}

type User struct {
	People
	username string
	password string
	email string
}

func main() {
	people := People{name: "Diego", lastname: "Almeida"}
	fmt.Println(people)

	user := User{
		People:   People{name: "Diego", lastname: "Almeida"},
		username: "dyhalmeida",
		password: "****",
		email: "dyhalmeida@gmail.com",
	}
	fmt.Println(user.name)
	fmt.Println(user.lastname)
	fmt.Println(user.username)
	fmt.Println(user.password)
	fmt.Println(user.email)
}