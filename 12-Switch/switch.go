package main

import "fmt"

func diaDaSemana(numero int) string{
	switch numero{
	case 1:
		return "Segunda-Feira"
	case 2:
		return "Terça-feira"
	case 3:
		return "Quarta-feira"
	case 4:
		return "Quinta-feira"
	case 5:
		return "Sexta-feira"
	case 6:
		return "Sábado"
	case 7:
		return "Domingo"
	default:
		return "Número Inválido"
	}
}

func diaDaSemana2(numero int) string{
	var diaDaSemana string
	switch {
	case numero == 1:
		diaDaSemana = "Segunda-Feira"
		fallthrough //joga a condição para a próxima comparação
	case numero == 2:
		diaDaSemana = "Terça-feira"
	case numero == 3:
		diaDaSemana = "Quarta-feira"
	case numero == 4:
		diaDaSemana = "Quinta-feira"
	case numero == 5:
		diaDaSemana = "Sexta-feira"
	case numero == 6:
		diaDaSemana = "Sábado"
	case numero == 7:
		diaDaSemana = "Domingo"
	default:
		diaDaSemana = "Número Inválido"
	}
	return diaDaSemana
}

func main() {
	fmt.Println("Switch")

	dia := diaDaSemana(3)
	fmt.Println(dia)

	fmt.Println("---------")
	dia2 := diaDaSemana2(4)
	fmt.Println(dia2)

	dia3 := diaDaSemana2(1)
	fmt.Println(dia3)

}