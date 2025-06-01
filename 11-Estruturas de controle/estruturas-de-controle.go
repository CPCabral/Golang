package main

import "fmt"

func main() {
	fmt.Println("Estruturas de Controle")

	numero := 10

	if numero > 15{
		fmt.Println("Maior que 15")
	} else{
		fmt.Println("Menor ou igual a 15")
	}

	//variável if init, existe apenas dentro do if
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Número é maior que zero")
	} else if outroNumero < 0 {
		fmt.Println("Número é menor que zero")
	} else {
		fmt.Println("Número é igual a zero")
	}
}