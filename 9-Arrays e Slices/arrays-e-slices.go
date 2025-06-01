package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("# Arrays e Slices")
	fmt.Println("# Arrays")

	var array1 [5]string
	array1[0] = "Posição 1"
	fmt.Println(array1)

	array2 := [5]string{"Posição 1","Posição 2", "Posição 3", "Posição 4", "Posição 5" }
	fmt.Println(array2)

	array3 := [...]int{1,2,3,4,5,6} //não é muito utilizado
	fmt.Println(array3)

	fmt.Println("# Slices")
	slice := []int{10, 11, 12, 13, 14}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice,18)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Posição alterada"
	fmt.Println(slice2)

	// ARRAYS INTERNOS
	fmt.Println("-----")
	slice3 := make([]float32, 10, 15) //cria um array e retorna um slice 
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //tamanho
	fmt.Println(cap(slice3)) //capacidade


}