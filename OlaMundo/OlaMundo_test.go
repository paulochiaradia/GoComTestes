package main

import "testing"

func testeOla(t *testing.T) {
	resultado := Ola()
	esperado := "Olá, !"

	if resultado != esperado {
		t.Errorf("Resultado '%s', Esperado '%s'", resultado, esperado)
	}
}
