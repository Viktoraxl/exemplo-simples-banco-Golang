package operacoes

import (
	"fmt"
)

type Operacao struct {
	Nome, Cpf     string
	ContaCorrente int
	Saldo         float64
}

func (c *Operacao) Sacar() string {
	var valor float64

	fmt.Println("Qual valor deseja sacar")
	fmt.Scan(&valor)

	if c.Saldo < valor {
		return "Saldo insuficiente"
	} else {
		c.Saldo -= valor
		return "Saque efetuado com sucesso."
	}
}

func (c *Operacao) Depositar() string {
	var valor float64

	fmt.Println("Qual valor deseja depositar?")
	fmt.Scan(&valor)
	c.Saldo = c.Saldo + valor

	return "Depósito efetuado com sucesso!"
}

func (c *Operacao) Transferir(contaDestino *Operacao) string {
	var valorDestino float64

	fmt.Println("Digite o valor da transferencia para o sr(a),", contaDestino.Nome, " : ")
	fmt.Scan(&valorDestino)
	if valorDestino < c.Saldo {
		c.Saldo -= valorDestino
		contaDestino.Saldo += valorDestino
		return "Transferencia realizada com sucesso."
	} else {
		return "Saldo Insuficiente"
	}
}

func (c *Operacao) Extrato() (string, float64) {
	return "Saldo atual: ", c.Saldo
}

func (c *Operacao) PagarBoleto() (string, float64) {
	var boleto float64
	fmt.Println("Qual o valor da fatura a ser paga?")
	fmt.Scan(&boleto)
	if c.Saldo < boleto {
		return "Saldo insuficiente: ", c.Saldo
	} else {
		c.Saldo -= boleto
		return "Fatura paga com sucesso. Novo saldo: ", c.Saldo
	}
}
