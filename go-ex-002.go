package main

import "fmt"

// Faça um programa que cria um vetor de inteiros com 10 elementos.
// Então inicialize este vetor com número quaisquer e imprima na tela todos os seus elementos.

func main() {
	vetor := [10]int{99, 88, 77, 66, 55, 44, 33, 22, 11}
	for i, v := range vetor {
		fmt.Println("Index", i, v)
	}
}
