package enderecos

import (
	"strings"
)
//TipoDeEndereco verifica se um endereço tem um tipo válido
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	enderecoEmLetraMinuscula := strings.ToLower(endereco)
	primeiraPalavraDoEndereco := strings.Split(enderecoEmLetraMinuscula, " ")[0] //transforma a entrada em um slice com o separador indicado

	enderecoTemTipoValido := false
	for _, tipo :=range tiposValidos{
		if tipo == primeiraPalavraDoEndereco{
			enderecoTemTipoValido = true
		}
	}
if enderecoTemTipoValido {
	return strings.Title(primeiraPalavraDoEndereco)
}

return "Tipo inválido"
}