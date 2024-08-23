package main

import "fmt"

// - Utilizando como base o exercício anterior, utilize slicing para demonstrar os valores:
// - Do primeiro ao terceiro item do slice (incluindo o terceiro item!)
// - Do quinto ao último item do slice (incluindo o último item!)
// - Do segundo ao sétimo item do slice (incluindo o sétimo item!)
// - Do terceiro ao penúltimo item do slice (incluindo o penúltimo item!)
// - Desafio: obtenha o mesmo resultado acima utilizando a função len() para determinar o penúltimo item

func main() {
	lista := []int{15, 62, 75, 84, 93, 100, 231, 31, 432, 320, 42}

	fmt.Printf("%v", lista[:3])
	fmt.Printf("%v", lista[4:])
	fmt.Printf("%v", lista[1:7])
	fmt.Printf("%v", lista[2:len(lista)-1])
}
