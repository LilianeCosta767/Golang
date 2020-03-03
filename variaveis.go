package main

import "fmt"

func main() {
	//como declarar variavel
	//valor padrao de variavel é 0
	var idade int = 13
	fmt.Println(idade)
	//outra forma de declarar uma variavel:
	var altura = 156 //se nao colocar o tipo precisa ser inicializada
	fmt.Println(altura)
	fmt.Println("Sua idade é ", idade, "e tem ", altura, " de altura")

	//declarando duas variaveis no mesmo comando 
	var ano, quantidade int = 1999, 28

	fmt.Println("\n\nano = ", ano, "\nquantidade = ", quantidade)

	//mais uma forma de se declarar variavel:
	var (
		nmCalcado int = 34
		nmMegaSena int = 65
		nomeCrush string = "Ian"
	)

	fmt.Println("\n\nnmCalcado = ", nmCalcado, "\nnmMegaSena = ", nmMegaSena, "\nnomeCrush = ", nomeCrush)

	matriculaAluno := 209918 //para nao colocar o "var" no inicio da declaração
	fmt.Println("\nmatriculaAluno = ", matriculaAluno)
}