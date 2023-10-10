package Iteracao

import "testing"

func TestRepeticoes(t *testing.T) {

	verificaMensagemCorreta := func(t *testing.T, esperado, resultado string) {
		t.Helper()
		if esperado != resultado {
			t.Errorf("Esperado %s - resultado %s", resultado, esperado)
		}
	}

	t.Run("Verifica Repete", func(t *testing.T) {
		resultado := Repetir("a", 5)
		esperado := "aaaaa"
		verificaMensagemCorreta(t, esperado, resultado)
	})
}

func BenchmarkRepetir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repetir("a", 5)
	}
}
