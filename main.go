package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

// Quando criamos uma função para a struct, temos que referenciar o obj -> (c *ContaCorrente)
// Assim a função fica disponivel para o obj instanciado
func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso."
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Depósito realizado com sucesso", c.saldo
	} else {
		return "Valor inválido", c.saldo
	}
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia < c.saldo && valorDaTransferencia > 0 {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}

}

func main() {

	conta1 := ContaCorrente{titular: "Silvia", saldo: 300}
	conta2 := ContaCorrente{titular: "Gustavo", saldo: 100}

	fmt.Println(conta1)
	fmt.Println(conta2)

	status := conta1.Transferir(200, &conta2)
	fmt.Println("Tranferência:", status)

	fmt.Println(conta1)
	fmt.Println(conta2)

}
