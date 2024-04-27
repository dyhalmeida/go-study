package main

import (
	"fmt" // Importa o pacote "fmt" para formatação e saída de dados
	"os"  // Importa o pacote "os" para operações relacionadas ao sistema operacional
)

func main() {
	// Define o nome do arquivo a ser criado
	filename := "example.txt"

	// Cria um novo arquivo com o nome especificado e abre para escrita
	file, err := os.Create(filename)
	// Verifica se ocorreu algum erro durante a criação do arquivo
	if err != nil {
		// Se houver um erro, o programa entra em estado de pânico e imprime o erro
		panic(err)
	}

	// Escreve os dados fornecidos (no formato de slice de bytes) no arquivo criado
	fileLen, err := file.Write([]byte("Exemplo de criação e escrita em arquivo"))
	// Verifica se ocorreu algum erro durante a escrita no arquivo
	if err != nil {
		// Se houver um erro, o programa entra em estado de pânico e imprime o erro
		panic(err)
	}

	// Imprime uma mensagem informando que o arquivo foi criado com sucesso
	fmt.Printf("O arquivo " + filename + " foi criado som sucesso! \n")
	// Imprime o tamanho do arquivo em bytes
	fmt.Printf("Tamanho do arquivo: %d bytes", fileLen)
}
