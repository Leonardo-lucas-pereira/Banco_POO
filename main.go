package main

import (
	"Banco/contas"
	"fmt"
)

func main() {
	contaExemplo := contas.ContaCorrente{}
	contaExemplo.Depositar(100)

	fmt.Println(contaExemplo)
	fmt.Println(contaExemplo.ObterSaldo())
}
