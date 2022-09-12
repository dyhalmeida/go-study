package main

import "fmt"

type User struct {
	name string
	lastname string
	age uint8
}

func (u User) getFullName() string {
	return u.name + " " + u.lastname
}

func (u User) isAdult() bool {
	return u.age >= 18
}

func (u *User) makeBirthday() {
	u.age++
}

func main() {

	user := User{"Diego", "Almeida", 30}
	fmt.Println(user.getFullName())
	fmt.Println(user.isAdult())
	fmt.Println(user.age)
	user.makeBirthday()
	fmt.Println(user.age)

}