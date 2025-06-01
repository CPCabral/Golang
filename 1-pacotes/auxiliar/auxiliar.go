package auxiliar

import "fmt"

// Escrever (Maiusculo- publico) se for minusculo fica privado, só pode ser acessada por arquivo no mesmo pacote

func Escrever() {
	fmt.Println("Escrevendo do arquivo auxiliar")
	escrever2()
}