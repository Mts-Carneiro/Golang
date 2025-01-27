package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos (n1, n2 int8) (int8, int8){
	soma := n1 + n2
	subtração := n1 - n2

	return soma, subtração
}

func main() {
	//Teste de função com unico retorno
	soma := somar(10, 20)
	fmt.Println(soma)

	//Teste de função com mais de um retorno
	calculo, calculo2 := calculosMatematicos(10, 20)
	fmt.Println(calculo, calculo2)

	//Teste de função com mais de um retorno, mas, retornando apenas o primeiro retorno
	resultado, _ := calculosMatematicos(10,15)
	fmt.Println(resultado)

}