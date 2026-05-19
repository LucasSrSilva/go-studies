package main

import (
	"fmt"

	t "poo/clientes"
	c "poo/contas"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	t1 := t.Titular{Nome: "roberto", CPF: "2123", Idade: 24}
	t2 := t.Titular{Nome: "silvio", Idade: 24}
	conta1 := c.ContaCorrente{Titular: t1, NumeroAgencia: 1, NumeroConta: 1234}
	conta2 := c.ContaCorrente{Titular: t2, NumeroAgencia: 1, NumeroConta: 1234}
	conta1.Depositar(500)
	conta2.Depositar(500)

	PagarBoleto(&conta1, 300)
	fmt.Println(conta1, conta2)
	conta1.Transferir(200, &conta2)
	fmt.Println(conta1, conta2)
}
