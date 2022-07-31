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
func (c *ContaCorrente) sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso."
	} else {
		return "Saldo insuficiente"
	}
}

func main() {

	conta := ContaCorrente{}
	conta.titular = "Leonardo"
	conta.numeroAgencia = 589
	conta.numeroConta = 123456
	conta.saldo = 20375.75

	fmt.Println(conta.saldo)
	fmt.Println(conta.sacar(60000))
	fmt.Println(conta.saldo)

}
