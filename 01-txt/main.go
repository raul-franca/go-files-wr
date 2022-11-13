package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//cria o arquivo ou trunca, caso não passe o path completo de onde o
	//arq deva ser cria/salvo ele sera criado na mesma pasta do go.mod
	file, err := os.Create("01-txt/arquivo.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//Write retorna o tamanho do arq len(b)
	tamanho, err := file.Write([]byte("Texto para o aquivo Raul"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("O  tamanho do arquivo é %d bytes\n", tamanho)

	// leitura de arq:
	//ReadFile() le o arq completo
	f, err := os.ReadFile("01-txt/arquivo.txt")
	if err != nil {
		panic(err)
	}
	println(string(f))

	//Leitor por partes
	f2, err := os.Open("01-txt/arquivo.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(f2) // retorna um leitor com o tamanho de buffer padrão
	buffer := make([]byte, 5)     // cria um slice de ‘bytes’ com tamanho de 3

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	err = os.Remove("01-txt/arquivo.txt")
	if err != nil {
		panic(err)
	}

}
