package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
	endereco endereco
}

type endereco struct{
	logradouro string
	numero uint8
}

func main() {
	var u usuario

	enderecoEx := endereco{"rua 1", 10}

	u.nome = "Matheus"
	u.idade = 21
	//Caso você instacie os campos, os campos nulos serão preenchidos com 0
	fmt.Println(u)

	//Tambem é possivel criar um struct dentro de outro struct
	u2 := usuario{"João", 22, enderecoEx} 

	fmt.Println(u2)
}