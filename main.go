package main

import (
	"Banco/clientes"
	"Banco/contas"
	"fmt"
)

func main() {

	cliente := clientes.Titular{Nome: "Leonardo", CPF: "123.456.895.56", Profissao: "Engenheiro"}
	fmt.Println(cliente)
	conta := contas.ContaCorrente{Titular: cliente, NumeroAgencia: 123, NumeroConta: 123456, Saldo: 89567.32}
	fmt.Println(conta)

}
