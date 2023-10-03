package main

// Refatoracao do codigo, mantem passando nos testes
const olaPortugues = "Olá, "
const olaEspanhol = "Hola, "
const olaFrances = "Bonjour, "

func Ola(nome string, idioma string) string {
	if nome == "" {
		nome = "mundo"
	}
	return prefixoDeSaudacao(idioma) + nome + "!"
}

func prefixoDeSaudacao(idioma string) (prefixo string) {
	switch idioma {
	case "espanhol":
		prefixo = olaEspanhol
	case "francês":
		prefixo = olaFrances
	default:
		prefixo = olaPortugues
	}
	return
}
func main() {
}
