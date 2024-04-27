package main

import (
	"fmt"      // Importa o pacote "fmt" para formatação e saída de dados
	"io"       // Importa o pacote "io" para operações de entrada e saída
	"net/http" // Importa o pacote "net/http" para realizar requisições HTTP
)

func main() {
	// Define a URL de onde os dados serão buscados
	url := "https://gist.githubusercontent.com/dyhalmeida/842f6bfb7c20c9c20de4c97d36b354ca/raw/37c0236a5261cb666b5656c4263158edd8063ca9/toggles.json"

	// Realiza uma requisição GET para a URL especificada e armazena a resposta em res
	res, err := http.Get(url)
	// Verifica se ocorreu algum erro durante a requisição
	if err != nil {
		// Se houver um erro, o programa entra em estado de pânico e imprime o erro
		panic(err)
	}
	// Garante que o corpo da resposta (res.Body) seja fechado após o término da execução da função main
	defer res.Body.Close()

	// Lê todo o conteúdo do corpo da resposta (res.Body) e armazena em data
	data, err := io.ReadAll(res.Body)
	// Verifica se ocorreu algum erro durante a leitura do corpo da resposta
	if err != nil {
		// Se houver um erro, o programa entra em estado de pânico e imprime o erro
		panic(err)
	}

	// Imprime os dados lidos do corpo da resposta na saída padrão
	fmt.Println(string(data))
	/** output
	  {
	      "toggles": [
	          {
	              "name": "myToggle",
	              "type": "release",
	              "value": false
	          }
	      ]
	  }
	**/
}
