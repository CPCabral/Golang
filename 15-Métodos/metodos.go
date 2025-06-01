package main

import "fmt"

type usuario struct {
	nome  string
	idade int16
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do usuário %s no banco de dados \n", u.nome)
}

func (u usuario) maiorDeIdade() {
	if u.idade >= 18{
		fmt.Printf("O usuário %s é maior de idade\n", u.nome)
	} else {
		fmt.Printf("O usuário %s é menor de idade\n", u.nome)
	}
}

func (u *usuario) fazerAniversario(){
	u.idade++
}

func main() {
	usuario1 := usuario{"Otávio", 1}
	usuario1.salvar()

	usuario2 := usuario{"Camila", 25}
	usuario2.salvar()

	usuario1.maiorDeIdade()

	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)
}