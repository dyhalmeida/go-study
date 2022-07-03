package main

import (
	"fmt"
	"pkg/myutils"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Hello Golang with packages")
	myutils.Sum()
	// myutils.avg() aqui não funciona por que avg não está sendo exportando com letra maiúscula
	error1 := checkmail.ValidateFormat("dyhalmeida@gmail.com")
	fmt.Println(error1)

	error2 := checkmail.ValidateFormat("dyhalmeida@.com")
	fmt.Println(error2)

	/*
		Output:
		Hello Golang with packages
		sum...
		avg...
		<nil>
		invalid format
	*/
}