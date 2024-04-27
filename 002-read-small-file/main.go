package main

import (
	"fmt" // Importa o pacote "fmt" para formatação e saída de dados
	"os"  // Importa o pacote "os" para operações relacionadas ao sistema operacional
)

func main() {
	// Define o nome do arquivo a ser lido
	filename := "example.txt"

	// Lê todo o conteúdo do arquivo especificado e armazena em "data"
	data, err := os.ReadFile(filename)
	// Verifica se ocorreu algum erro durante a leitura do arquivo
	if err != nil {
		// Se houver um erro, o programa entra em estado de pânico e imprime o erro
		panic(err)
	}

	// Imprime na saída padrão os dados lidos do arquivo convertidos para string
	fmt.Println(string(data))
}
