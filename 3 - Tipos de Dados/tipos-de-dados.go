package main

import (
	"errors"
	"fmt"
)

func main(){
	var numero int16 = 100
	fmt.Println(numero)

	var numeroReal1 float32 = 1233.25
	fmt.Println(numeroReal1)

	var boleano bool
	fmt.Println(boleano)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}

