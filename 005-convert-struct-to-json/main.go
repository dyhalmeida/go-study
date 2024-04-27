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
	// Cria uma instância da estrutura User com valores predefinidos
	user := User{
		Name:     "Diego Almeida",          // Define o nome do usuário
		Username: "dyhalmeida",             // Define o nome de usuário
		Email:    "diego.almeida@mail.com", // Define o endereço de email do usuário
	}

	// Serializa a estrutura user para formato JSON e armazena o resultado na variável jsonData
	jsonData, err := json.Marshal(user)
	// Verifica se ocorreu algum erro durante a serialização
	if err != nil {
		// Se houver um erro, o programa entra em estado de pânico e imprime o erro
		panic(err)
	}

	// Imprime o JSON serializado na saída padrão
	fmt.Println(string(jsonData))
	/** output:
	  {"name":"Diego Almeida","username":"dyhalmeida","email":"diego.almeida@mail.com"}
	  **/
}
