package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Loops")

	i := 0
	 for i < 10 {
		i++
		fmt.Println("Incrementando i:", i)
		time.Sleep(time.Second)
	 }
	 fmt.Println(i)
	 fmt.Println("-------------")

	 for j := 0; j < 10; j += 2 {
		fmt.Println("Incrementando j:", j)
		time.Sleep(time.Second)
	 }
	 //fmt.Println(j) //a variável só existe dentro do for

	 //for range
	 nomes := [3]string{"João", "Lucas", "Mateus"}

	 for indice, nome := range nomes{
		fmt.Println(indice, nome)
	 }
fmt.Println("---------------")
	 	 for _, nome := range nomes{
		fmt.Println(nome)
	 }
fmt.Println("---------------")

for i, letra := range "PALAVRA" {
	fmt.Println(i, letra)	
}
fmt.Println("---------------")
for i, letra := range "PALAVRA" {
	fmt.Println(i, string(letra))	
}

usuario := map[string]string{
	"nome" : "Leonardo",
	"sobrenome" : "Silva",
}

for chave, valor := range usuario{
	fmt.Println(chave, valor)
}

//loop infinito
for {
	fmt.Println("Executando infinitamente")
	time.Sleep(time.Second)
}

//Não pode iterar em struct
}