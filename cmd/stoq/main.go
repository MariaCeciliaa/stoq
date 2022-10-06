package main

import "fmt"

type Pessoa struct {
	nome string
	idade string
	telefone string 
	email string
}

func (p Pessoa) Falar() {
	fmt.Printf("Olá eu sou %v, tenho %v anos, meu e-mail é: %v\n", p.nome, p.idade, p.email)
}

func (p Pessoa) Comer() {
	fmt.Printf("Estou comendo...")
}

func main() {

	pessoa1 := Pessoa{
		nome: "Cecilia",
		idade: "19",
		telefone: "9658745874",
		email: "cecilia@email.com",
	}

	pessoa2 := Pessoa{
		nome: "Amanda",
		idade: "15",
		telefone: "9658745874",
		email: "amanda@email.com",
	}

	pessoa1.Falar()
	pessoa2.Falar()

	pessoa1.Comer()
}

