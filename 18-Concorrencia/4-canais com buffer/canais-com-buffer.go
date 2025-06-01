package main

import "fmt"

//o canal deve ser aberto em uma função diferente, pois o canal bloqueia quando recebe um valor
//para abrir na mesma função deve aplicar um buffer no canal, só bloqueia quando atingir a capacidade maxima determinada

func main(){
	canal := make(chan string, 2) //canal com buffer
	canal <- "olá mundo!"
	canal <- "Programando em GO"

	mensagem := <- canal
	mensagem2 := <- canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}