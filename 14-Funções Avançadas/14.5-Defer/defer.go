package main

import "fmt"

func funcao1(){
	fmt.Println("Executando a função 1")
}

func funcao2(){
	fmt.Println("Executando a função 2")
}

func alunoAprovado(n1, n2 float32) bool{
	defer fmt.Println("Média Calculada. Resultado será retornado") // vai ser executado imediatamente antes do return
	fmt.Println("Entrando na função para verificar se o aluno será aprovado")
	media := (n1 + n2)/2

	if media >= 6 {
		return true
	}
	return false
}

func main() {
	fmt.Println("Defer") //adia a execução de uma função

	defer funcao1()
	funcao2()
	fmt.Println(alunoAprovado(7,8))
}