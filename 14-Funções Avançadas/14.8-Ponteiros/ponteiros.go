package main

import "fmt"

func inverterSinal(numero int) int {
	return numero * -1
}

func inverterSinalcomPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	numero := 20 // não mexe no valor da variavel, faz uma cópia
	fmt.Println("número invertido: ", inverterSinal(numero))
	fmt.Println("numero:", numero)

	fmt.Println("-------------")
	novoNumero := 40 //com ponteiro altera a variavel armazenada, está passando a referencia da variavel
	fmt.Println("numero:", novoNumero)
	inverterSinalcomPonteiro(&novoNumero)
	fmt.Println("numero invertido: ", novoNumero )


}