package main

import (
	"fmt"
)

// Refatoracao do codigo, mantem passando nos testes
const olaPortugues = "Olá, "
const olaEspanhol = "Hola, "

func Ola(nome string, idioma string) string {
	if nome == "" {
		nome = "mundo"
	}
	switch idioma {
	case "espanhol":
		return olaEspanhol + nome + "!"
	default:
		return olaPortugues + nome + "!"
	}
}
func main() {
	fmt.Println(Ola("Chris", ""))
}
