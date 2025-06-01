package main

import (
	"errors"
	"fmt"
)

func main() {
	//int8, int16, int32, int64
	var numero int16 = 10000
	fmt.Println(numero)

	var numero2 int = -1000000000000000000
	fmt.Println(numero2)
	
	numero3 := 1000000000000000000
	fmt.Println(numero3)
	
	//uint numero sem sinal

	//alias
		//int32 = rune
	var numero4 rune = 12345
	fmt.Println(numero4)
		
		//uint8 = byte
	var numero5 byte = 123
	fmt.Println(numero5)
		
	// float32, float64
	var numeroReal float32 = 123.5

	numeroReal2 := 1234.56
	fmt.Println(numeroReal,numeroReal2)

	//Strings
	var str string = "texto"
	fmt.Println(str)

	str2 := "Texto 2"
	fmt.Println(str2)

	char := 'B' // ' ' retorna posição o que o caractere esta na tabela ascii
	fmt.Println(char)
	//fim strings

	var texto string
	fmt.Println(texto)

	var texto2 int16
	fmt.Println(texto2)

	var booleano bool = true
	fmt.Println(booleano)

	var booleano1 bool
	fmt.Println(booleano1)

	var erro error
	fmt.Println(erro)

	var erro1 error = errors.New("Erro interno")
	fmt.Println(erro1)
}