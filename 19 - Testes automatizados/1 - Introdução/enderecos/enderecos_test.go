package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

//Funções de teste iniciam com Test
func TestTipoDeEndereco(t *testing.T) {
	t.Parallel()

	cenarioDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Rodovia dos imigrantes", "Rodovia"},
		{"Avenida Paulista", "Avenida"},
		{"Praça das rosas", "Tipo Inválido"},
		{"Estrada ABC", "Estrada"},
		{"RUA DOS BOBOS", "Rua"},
		{"rua dos bobos", "Rua"},
		{" ", "Tipo Inválido"},
	}

	for _, cenario := range cenarioDeTeste {
		tipoEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if tipoEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e recebeu %s",
				tipoEnderecoRecebido,
				cenario.retornoEsperado,
			)
		}
	}

	// enderecoParaTeste := "Rua Paulista"
	// tipoDeEnderecoEsperado := "Avenida"
	// tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)

	// if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
	// 	t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e recebeu %s",
	// 		tipoDeEnderecoEsperado,
	// 		tipoDeEnderecoRecebido,
	// 	)
	// }
}

func TestQualquer(t *testing.T) {
	t.Parallel()

	if 1 > 2 {
		t.Errorf("Teste Quebrou!")
	}
}
