package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")
	var variavel1 int = 10
	var variavel2 int = variavel1 // cópia

	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println(variavel1, variavel2)

	//Ponteiro é uma referência de memoria
	var variavel3 int
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3

	fmt.Println(variavel3, ponteiro)
	fmt.Println(variavel3, *ponteiro) //desreferenciação

}