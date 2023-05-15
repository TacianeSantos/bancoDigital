package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

type VerificarConta interface {
	Sacar(float64) string
}

func PagarBoleto(conta VerificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

func main() {

	clienteTaciane := clientes.Titular{Nome: "Taciane", CPF: "123.123.123-00", Profissao: "Desenvolvedora"}

	novaConta := contas.ContaCorrente{Titular: clienteTaciane, Agencia: 123, ContaCorrente: 1234569}

	mensagem, saldo := novaConta.Deposito(1000)

	fmt.Println(mensagem, saldo)

	fmt.Println(novaConta.ObterSaldo())

	PagarBoleto(&novaConta, 50)

	fmt.Println(novaConta.ObterSaldo())
}
