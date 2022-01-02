package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Digite seu nome: ")
	var nome string
	fmt.Scanln(&nome)
	nome = strings.TrimSpace(nome)
	fmt.Printf("Oi, %s! Eu sou GO", nome)
}
