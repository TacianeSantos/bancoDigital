package contas

import (
	cliente "banco/clientes"
	"fmt"
)

type ContaCorrente struct {
	Titular       cliente.Titular
	Agencia       int
	ContaCorrente int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorDoSaque
		fmt.Println(c.saldo)

		return "Saque realizado com sucesso"
	} else {
		return "saldo insulficiente"
	}
}

func (c *ContaCorrente) Deposito(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso", c.saldo
	} else {
		return "Operação não realizada", c.saldo
	}
}

func (c *ContaCorrente) Transferir(valorDaTranferencia float64, contaCorrenteDestino *ContaCorrente) bool {
	if valorDaTranferencia > 0 && c.saldo >= valorDaTranferencia {
		c.Sacar(valorDaTranferencia)
		contaCorrenteDestino.Deposito(valorDaTranferencia)
		return true
	} else {
		return false
	}
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
