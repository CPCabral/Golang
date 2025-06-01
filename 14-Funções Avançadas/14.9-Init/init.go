package main

import "fmt"

var n int
func main() {
fmt.Println("Inicializando a função main.")
fmt.Println(n)
}

func init() { //inicializa antes de qualquer função
	fmt.Println("Inicializando a função Init.")
	n = 10
}