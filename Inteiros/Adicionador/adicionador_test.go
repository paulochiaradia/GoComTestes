package main

import "testing"

func TestAdicionador(t *testing.T) {

	verificaMensagemCorreta := func(t *testing.T, resultado, esperado int) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("Resultado '%d',\n Esperado '%d'", resultado, esperado)
		}
	}

	t.Run("Adiciona 2 + 2", func(t *testing.T) {
		soma := adiciona(2, 2)
		esperado := 4
		verificaMensagemCorreta(t, soma, esperado)
	})
}
