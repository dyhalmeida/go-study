package main

import (
	"encoding/json" // Importa o pacote "encoding/json" para manipulação de JSON
	"fmt"           // Importa o pacote "fmt" para formatação e saída de dados
)

// Define a estrutura "User" que representa um usuário com campos Name, Username e Email
type User struct {
	Name     string `json:"name"`     // Campo "Name" da estrutura User com a tag json "name"
	Username string `json:"username"` // Campo "Username" da estrutura User com a tag json "username"
	Email    string `json:"email"`    // Campo "Email" da estrutura User com a tag json "email"
}

func main() {
	// Define uma variável jsonData que armazena os dados JSON como um slice de bytes
	jsonData := []byte(`{"name":"Diego Almeida","username":"dyhalmeida","email":"diego.almeida@mail.com"}`)

	// Declara uma variável user do tipo User
	var user User

	// Deserializa o JSON contido em jsonData para a estrutura user
	err := json.Unmarshal(jsonData, &user)
	// Verifica se houve algum erro na desserialização do JSON
	if err != nil {
		// Se houver um erro, o programa entra em estado de pânico e imprime o erro
		panic(err)
	}

	// Imprime os campos da estrutura User
	fmt.Println(user.Name)     // Imprime o nome do usuário
	fmt.Println(user.Username) // Imprime o nome de usuário
	fmt.Println(user.Email)    // Imprime o endereço de email do usuário
}
