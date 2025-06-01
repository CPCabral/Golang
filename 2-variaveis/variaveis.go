package main

// formas de declarar variaveis
import "fmt"

func main() {
	var variavel1 string = "variavel 1"
	variavel2 := "variavel 2"
	fmt.Println(variavel1,variavel2)
	var (
		variavel3 string = "akjfbh"
		variavel4 string = "djnfkade"
	)
	fmt.Println(variavel3)
	fmt.Println(variavel4)

	variavel5, variavel6 := "variavel 5", "variavel 6"
	fmt.Println(variavel5,variavel6)

	const constante1 string = "constante 1"
	fmt.Println(constante1)

	// inverter as variaveis
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5,variavel6)
}