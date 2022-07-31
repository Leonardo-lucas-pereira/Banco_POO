package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {

	conta := ContaCorrente{titular: "Leonardo",
		numeroAgencia: 589, numeroConta: 123456, saldo: 20375.75}

	fmt.Println(conta)

	conta2 := ContaCorrente{"Lucas", 456, 654321, 897.45}

	fmt.Println(conta2)
}
