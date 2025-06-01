package main

import "fmt"

func main() {
	// 	ARITIMETICOS
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 3

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	// var numero1 int16 = 10
	// var numero2 int32 = 25
	// soma2 := numero1 + int16(numero2) não pode fazer operações com numeros de formatos diferentes

	// FIM DOS ARITMETICOS

	// ATRIBUIÇÃO
	 var variavel1 string = "string"
	 variavel2 := "string"
	 fmt.Println(variavel1, variavel2)
	// FIM ATRIBUIÇÃO

	// OPERADORES RELACIONAIS
	fmt.Println(1 > 2, 1 >= 2, 1 == 2, 1 <= 2, 1 < 2, 1 != 2)
	//FIM DOS RELACIONAIS

	//OPERADORES LOGICOS
	fmt.Println("---------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso) //and
	fmt.Println(verdadeiro || falso) //ou
	fmt.Println(!verdadeiro)
	//FIM OPERADORES LOGICOS

	//OPERADORES UNÁRIOS
	numero := 10
	numero ++ // numero = numero + 1
	fmt.Println(numero)
	numero += 15 // numero = numero + 15
	fmt.Println(numero)
	numero --
	numero -= 10 //numero = numero - 10
	numero *= 3 // numero = numero * 3
	numero /= 10
	numero %=3
	// FIM DOS OPERADORES UNARIOS

	// OPERADORES TERNARIOS = não existem no go, usa ir

	var texto string
	if numero > 5{
		texto = "maior que 5"
	} else{
		texto = "menor que 5"
	}
	fmt.Println(texto)


}