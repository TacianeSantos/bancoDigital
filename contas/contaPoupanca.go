package contas

import (
	"banco/clientes"
	"fmt"
)

type ContaPoupanca struct {
	Titular                                      clientes.Titular
	NumeroAgencia, NumeroContaPoupanca, OPeracao int
	saldo                                        float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorDoSaque
		fmt.Println(c.saldo)

		return "Saque realizado com sucesso"
	} else {
		return "saldo insulficiente"
	}
}

func (c *ContaPoupanca) Deposito(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso", c.saldo
	} else {
		return "Operação não realizada", c.saldo
	}
}
