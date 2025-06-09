package enderecos // ou usar enterecos_test (daí precisa importar o pacote)

import "testing"
// . "intoducao-testes/enderecos" (importa o pacote da função que vai ser testada como Alias)

// comando terminal vai entrar nos aquitos que tem "_test"
// TESTE DE UNIDADE

type cenarioDeTeste struct{
	enderecoInserido string
	retornoEsperado string
}

func TestTipoDeEndereco(t *testing.T) {
	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		// {"Praça das Arvores", "Tipo inválido"},
		{"RUA ABC", "Rua"},
		// {"", "Tipo inválido"},
	}

	for _, cenario := range cenariosDeTeste{
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if tipoDeEnderecoRecebido != cenario.retornoEsperado{
			t.Errorf("O tipo recebido %s é diferente do esperado %s",
				tipoDeEnderecoRecebido,
				cenario.retornoEsperado,
			)
		}
	}
}

// func TestTipoDeEndereco(t *testing.T) {
// 	enderecoParaTeste := "Avenida Paulista"	
// 	tipoDeEnderecoEsperado := "Avenida"
// 	tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)

// 	if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado{
// 		t.Errorf("O tipo recebido é diferente do esperado. Esperava %s e recebeu %s",
// 			tipoDeEnderecoEsperado,
// 			tipoDeEnderecoRecebido,
// 		)
// 	}
// }

// no terminal: go test
// go test ./... (entra em todas as pastas)
// go test -v (traz mais informações)
// t.Parallel() //em todas as funções para elas rodarem em paralelo

// go test --cover (cobertura do teste)
// go test --coverprofile nomedoarquivo.tx (cria arquivo txt com os dados da cobertura do teste)
// go tool cover --func=nomedoarquivo.txt (le o arquivo txt no terminal, traz poucas informações)
// go tool cover --html=nomedoarquivo.txt (cria um html detalhado do que está ou não sendo coberto pelo teste)