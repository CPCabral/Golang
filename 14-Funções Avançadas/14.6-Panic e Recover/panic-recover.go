package main

import "fmt"
func recuperarExecução(){
	fmt.Println("Tentativa de recuperar a execução")
	if r := recover(); r != nil{
		fmt.Println("Execução recuperada com sucesso!")
	}
}

func consultaAlunoAprovado(n1,n2 float64) bool{
	defer recuperarExecução()
	media := (n1 + n2)/2

	if media > 6 {
		return true
	}else if media < 6 {
		return false
	}
	panic("A média é exatamente 6") //Como não tem solução o programa para por não saber o que fazer se não for tratado, antes de matar o programa ele chama as funções que tem defer
}

func main() {
	fmt.Println(consultaAlunoAprovado(6,6))
	fmt.Println("Pós execução")

}