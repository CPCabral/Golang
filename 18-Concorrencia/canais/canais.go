package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go escrever("Olá mundo!", canal)

	fmt.Println("depois da função escrever começar a ser executada")
	
	// for {
	// mensagem, aberto := <- canal //ele vai esperar receber um valor	
	// if !aberto{
	// 	break
	// }
	// fmt.Println(mensagem)
	// }

	for mensagem := range canal{
		fmt.Println(mensagem)
	}
	fmt.Println("fim do programa!")
}

func escrever(texto string, canal chan string) {
	// time.Sleep(time.Second * 5)
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		canal <- texto
		time.Sleep(time.Second)
	}
	close(canal)
}

//deadlock -> quando não tem nenhum canal enviando dados mas tem um canal esperando receber dados
// !!! só pega em execução, não pega em compilação