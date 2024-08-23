package main

import "fmt"

// - Usando uma literal composta:
// - Crie um array que suporte 5 valores to tipo int
// - Atribua valores aos seus índices
// - Utilize range e demonstre os valores do array.
// - Utilizando format printing, demonstre o tipo do array.

func main() {
	lista := [5]int{15, 62, 75, 84, 93}

	for i, v := range lista {
		println(i, v)
	}
	fmt.Printf("%T", lista)
}
