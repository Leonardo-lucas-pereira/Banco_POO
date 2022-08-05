package main

import (
	"Banco/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	conta1 := contas.ContaCorrente{}
	conta1.Depositar(100)
	conta1.Sacar(80)
	fmt.Println(conta1)
	fmt.Println(conta1.ObterSaldo())
	PagarBoleto(&conta1, 10)
	fmt.Println(conta1.ObterSaldo())
}
