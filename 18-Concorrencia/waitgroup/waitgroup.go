package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Sincronizar as goroutines 
	var waitGroup sync.WaitGroup

	waitGroup.Add(4)

	go func(){
		escrever("Olá mundo")
		waitGroup.Done() //-1
	}()

	go func(){
		escrever("Programando em GO")
		waitGroup.Done() //-1
	}()

		go func(){
		escrever("Olá mundo 2")
		waitGroup.Done() //-1
	}()

	go func(){
		escrever("Programando em GO 2")
		waitGroup.Done() //-1
	}()

	waitGroup.Wait() //0
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}