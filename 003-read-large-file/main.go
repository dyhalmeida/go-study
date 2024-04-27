package main

import (
	"bufio" // Importa o pacote "bufio" para leitura de arquivos em modo bufferizado
	"fmt"   // Importa o pacote "fmt" para formatação e saída de dados
	"os"    // Importa o pacote "os" para operações relacionadas ao sistema operacional
	"time"  // Importa o pacote "time" para manipulação de tempo
)

func main() {
	// Define o nome do arquivo a ser lido
	filename := "example.txt"

	// Abre o arquivo especificado em modo leitura e armazena o descritor do arquivo em "file"
	file, err := os.Open(filename)
	// Verifica se ocorreu algum erro ao abrir o arquivo
	if err != nil {
		// Se houver um erro, o programa entra em estado de pânico e imprime o erro
		panic(err)
	}

	// Cria um novo leitor de bufio associado ao arquivo aberto
	reader := bufio.NewReader(file)
	// Define um slice de bytes "chunck" com tamanho 2 para armazenar os dados lidos do arquivo em cada iteração

	chunck := make([]byte, 2)

	// Loop infinito para ler o arquivo até o final
	for {
		// Lê até len(chunck) bytes do arquivo para o slice "chunck" e armazena o número de bytes lidos em "data"
		data, err := reader.Read(chunck)
		// Verifica se ocorreu algum erro durante a leitura do arquivo
		if err != nil {
			// Se houver um erro, o loop é interrompido
			break
		}
		// Pausa a execução do programa por 1 segundo
		time.Sleep(time.Second)
		// Imprime na saída padrão os dados lidos do arquivo convertidos para string
		fmt.Print(string(chunck[:data]))
	}
}
