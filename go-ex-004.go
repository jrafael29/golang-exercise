package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// Faça um programa que abra um arquivo texto "input.txt", pré-existente. O programa então deve ler o arquivo linha por linha e apresentar a soma total do comprimento de todas as linhas. Verifique se é necessário fechar o arquivo antes do programa terminar. Adicione também suporte para eventuais situações de erro, como por exemplo não conseguir abrir o nome de arquivo especificado.

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Arquivo nao encontrado")
	}
	arr := strings.SplitAfter(string(file), "\n")

	for i, v := range arr {
		fmt.Printf("A linha %v Possui comprimento de %v \n", i, len(v))
	}
}
