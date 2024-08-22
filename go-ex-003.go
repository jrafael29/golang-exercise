package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

// Faça uma função/método que receba uma string como parâmetro e que retorne uma nova string, onde a sequência dos caracteres foi invertida. Dentro da parte principal (main), leia uma string digitada pelo usuário e passe para a função/método criada, imprimindo em seguida a string devolvida.

func revertString(str string) (string, error) {
	if len(str) > 0 {
		var reversedString string

		var reversedStrArr []string
		vetor := strings.SplitAfter(str, "")

		for i := len(vetor); i > 0; i-- {
			reversedStrArr = append(reversedStrArr, vetor[i-1])
		}
		reversedString = strings.Join(reversedStrArr, reversedString)
		return reversedString, nil
	}
	return "", errors.New("invalid input")
}

func main() {
	var strInput string
	fmt.Println("Informe a palavra que deseja inverter")
	fmt.Scanln(&strInput)

	reversedString, err := revertString(strInput)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reversedString)
}
