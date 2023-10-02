package main

import "fmt"

// Refatoracao do codigo, mantem passando nos testes
const olaPortugues = "Olá, "

func Ola(nome string) string {
	if nome == "" {
		nome = "mundo"
	}
	return olaPortugues + nome + "!"
}
func main() {
	fmt.Println(Ola("Chris"))
}
