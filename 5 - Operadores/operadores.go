package main

import "fmt"

func main() {
	//Aritimetricos
	soma := 1 + 2
	sutracao := 1 - 2
	divisao := 4 / 2
	multiplicacao := 2 * 2
	restoDivisao := 10 % 2

	fmt.Println(soma, sutracao, multiplicacao, divisao,restoDivisao)
	
	//Só haverá operação entre dois numeros se ambos forem da mesma tipagem

	var numero1 int16 = 10
	var numero2 int16 = 15
	soma2 := numero1 + numero2

	fmt.Println(soma2)


	//Atribuição

	var variavel1 string = "string"
	variavel := "String"

	println(variavel, variavel1)


	//Relacionais
	//Relacionais sempre retornam Boleanos
	fmt.Println(1>2)
	fmt.Println(1>=2)
	fmt.Println(1==2)
	fmt.Println(1!=2)
	
	verdadeiro, falso := true, false

	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)

	//Unários

	numero :=10
	numero++

	fmt.Println(numero)
	
	numero +=15
	
	fmt.Println(numero)
	
	numero--

	fmt.Println(numero)


	//O Go não possue operadores ternarios, para ultilizar algo parecido é necessario usar o If/Else
}