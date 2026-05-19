package contas

import c "poo/clientes"

type ContaCorrente struct {
	Titular       c.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valor float64) string {
	podeSacar := valor <= c.saldo && valor > 0
	if podeSacar {
		c.saldo -= valor
		return "Saque realizado"
	}
	return "Não foi possível sacar"
}
func (c *ContaCorrente) Depositar(valor float64) string {
	podeDepositar := valor > 0
	if podeDepositar {
		c.saldo += valor
		return "Deposito realizado"
	}
	return "Não foi possível depositar"
}

func (c *ContaCorrente) Transferir(valor float64, destino *ContaCorrente) string {
	podeTransferir := valor <= c.saldo && valor > 0
	if podeTransferir {
		c.Sacar(valor)
		destino.Depositar(valor)
		return "Transferencia realizada"
	}
	return "Não foi possível transferir"
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
