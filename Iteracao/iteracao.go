package Iteracao

const quantidadeRepeticoes = 5

func Repetir(caractere string, repeticoes int) string {
	var stringRepetida string
	for i := 0; i < repeticoes; i++ {
		stringRepetida += caractere
	}
	return stringRepetida
}
