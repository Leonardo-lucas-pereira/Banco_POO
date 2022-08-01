package main

import (
	"Banco/clientes"
	"Banco/contas"
	"fmt"
)

func main() {

	cliente1 := clientes.Titular("Leonardo", "123.456.871.97", "Engenheiro de Dados")
	conta1 := contas.ContaCorrente(cliente1, 123, 123456, 89568.45)

	fmt.Println(cliente1)
	fmt.Println(conta1)

}
