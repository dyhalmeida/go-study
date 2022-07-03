package myutils

import "fmt"

// Uma função de um pacote só ode ser expotada se estiver com a primeira letra maiúscula
// porém funções com letras minúsculas ou maiúsculas de outros arquivos do mesmo pacote,
// podem ser importando em arquivos do mesmo pacote
func Sum() {
	fmt.Println("sum...")
	avg() // aqui funciona por que faz parte do mesmo pacote
}