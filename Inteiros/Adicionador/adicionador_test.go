package inteiros

import (
	"fmt"
	"testing"
)

func TestAdicionador(t *testing.T) {

	verificaMensagemCorreta := func(t *testing.T, resultado, esperado int) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("Resultado %d - Esperado %d", resultado, esperado)
		}
	}

	t.Run("Adiciona 2 + 2", func(t *testing.T) {
		resultado := Adiciona(2, 2)
		esperado := 4
		verificaMensagemCorreta(t, resultado, esperado)
	})
}
func ExampleAdiciona() {
	soma := Adiciona(1, 5)
	fmt.Println(soma)
	// Output: 6
}
