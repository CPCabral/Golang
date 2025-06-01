package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
	endereco endereco
}
type endereco struct {
	logradouro string
	numero uint8
}

func main() {
	var u usuario
	u.nome = "Davi"
	u.idade = 21
	fmt.Println(u)

	enderecoExemplo := endereco{"rua dos bobos", 0}

	usuario2 := usuario{"Otavio", 1, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Camila"}
	fmt.Println(usuario3)
}