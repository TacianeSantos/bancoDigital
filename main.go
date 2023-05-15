package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func main() {

	clienteTaciane := clientes.Titular{Nome: "Taciane", CPF: "123.123.123-00", Profissao: "Desenvolvedora"}

	novaConta := contas.ContaCorrente{Titular: clienteTaciane, Agencia: 123, ContaCorrente: 1234569}

	mensagem, saldo := novaConta.Deposito(1000)

	fmt.Println(mensagem, saldo)

	fmt.Println(novaConta.ObterSaldo())
}
