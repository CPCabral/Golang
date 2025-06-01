package main

import "fmt"

func soma(numeros ... int) int {
	total := 0
	for _, numero := range numeros{
		total += numero
	}
	return total
}

func escrever(texto string, numero ... int){ //só pode ter uma variavel variatica e esta deve ser a ultima 
	for _, num := range numero{
		fmt.Println(texto, num)
	}
}
func main() {
	fmt.Println(soma(1,2,3,4,5,200))
	escrever("Camila", 1, 2, 3, 4, 19, 47)
}