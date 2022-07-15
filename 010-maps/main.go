package main

import "fmt"

func main(){
	maps()
}

func maps(){
	/*
		Um map em Go, é uma estrutura de dados imutável, no qual
		CHAVE e VALOR devem seguir a tipagem dedefinida para o map.
		Se um map for definido como chave do tipo string e valor do tipo string,
		TODAS as chaves e valores devem ser string.
	*/

	/* Aqui inicializo um map*/
	user := map[string]string{
		"name": "Diego",
		"lastname": "Almeida",
	}
	fmt.Println(user)

	/*Aqui eu obtenho o valor de uma chave (não acessível com notação de ponto(.)*/
	fmt.Println(user["name"])

	/*Aqui eu deleto uma chave do mao*/
	delete(user, "lastname")
	fmt.Println(user)

	/*Aqui eu adiciono uma chave no map, respeitando o tipo do valor*/
	user["city"] = "Camaçari-BA"
	fmt.Println(user)

	/*Aqui eu inicializo um map aninhado*/
	address := map[string]map[string]string{
		"street": {
			"name": "Rua das palmas",
			"number": "07",
		},
		"zip": {
			"code": "42808139",
		},
	}
	fmt.Println(address)

	/*Aqui eu removo uma chave do map aninhado*/
	delete(address, "zip")
	fmt.Println(address)

	/*Aqui eu adiciono uma chave no map aninhando*/
	address["zip"] = map[string]string{
		"code": "48280000",
	}
	fmt.Println(address)

}