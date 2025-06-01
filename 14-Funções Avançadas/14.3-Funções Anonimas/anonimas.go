package main

import "fmt"

func main() {
	// func() {
	// 	fmt.Println("olá mundo")
	// }()

	func(texto string) {
		fmt.Println(texto)
	}("olá mundo")

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto) // não printa na tela, só gera uma variável
	}("Passando Parâmetro")

	fmt.Println(retorno)
	
	const name, age = "Kim", 22
	fmt.Print(fmt.Sprintf("%s is %d years old.\n", name, age))

variavel := "Valor Qualquer"
fmt.Println("Valor da minha varíavel: " + variavel) // Vai escrever na tela
fmt.Printf("Valor da minha varíavel: %s\n", variavel) // Vai escrever na tela
variavel2 := fmt.Sprintf("Valor da minha variável: %s", variavel) // Vai retornar uma nova variável
fmt.Println(variavel2)
}