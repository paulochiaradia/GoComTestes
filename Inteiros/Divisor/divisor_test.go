package main

import "testing"

func TestDivisor(t *testing.T) {
	verificaMensagemCorreta := func(t *testing.T, resultado, esperado int) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("Resultado: %d \n Esperado: %d", resultado, esperado)
		}
	}

	t.Run("Divide 2 / 2", func(t *testing.T) {
		esperado := 1
		resultado := divide(2, 2)
		verificaMensagemCorreta(t, resultado, esperado)
	})
}
