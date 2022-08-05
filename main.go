package main

import (
	"Banco/contas"
	"fmt"
)

func main() {
	contaPoupanca := contas.ContaPoupanca{}
	contaPoupanca.Depositar(100)
	contaPoupanca.Sacar(80)
	fmt.Println(contaPoupanca)
	fmt.Println(contaPoupanca.ObterSaldo())
}
