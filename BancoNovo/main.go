package main

import (
	"fmt"
	"os"

	operacoes "github.com/victor/Treinamentos/BancoNovo/Operacoes"
)

func main() {
	var opcao, contaDestino int
	contaJoao, contaMaria, contaJose := operacoes.Operacao{Nome: "Joao", ContaCorrente: 1112223, Saldo: 1250.50},
		operacoes.Operacao{Nome: "Maria", ContaCorrente: 2223334, Saldo: 870.50},
		operacoes.Operacao{Nome: "Jose", ContaCorrente: 3334445, Saldo: 410.29}

	contas := []*operacoes.Operacao{&contaJoao, &contaMaria, &contaJose}

	for {
		cc := DigiteConta()
		var operacaoRealizada bool
		for {
			for _, conta := range contas {
				if conta.ContaCorrente == cc {
					fmt.Println("Bem vindo ", conta.Nome, ". O que deseja fazer?")
					fmt.Println("Saldo: ", conta.Saldo)
					fmt.Println("1 - Sacar")
					fmt.Println("2 - Depositar")
					fmt.Println("3 - Transferir")
					fmt.Println("4 - Extrato")
					fmt.Println("5 - Pagar uma conta")
					fmt.Println("6 - Sair")
					fmt.Scan(&opcao)

					switch opcao {
					case 1:
						fmt.Println(conta.Sacar())
						operacaoRealizada = true
					case 2:
						fmt.Println(conta.Depositar())
						operacaoRealizada = true
					case 3:
						fmt.Println("Digite a conta destino")
						fmt.Scan(&contaDestino)
						for _, contaD := range contas {
							if contaDestino == contaD.ContaCorrente {
								fmt.Println(conta.Transferir(contaD))
								fmt.Println("Sr(a)", conta.Nome, " seu saldo é: ", conta.Saldo)
								fmt.Println("O saldo do sr(a)", contaD.Nome, " é: ", contaD.Saldo)
								operacaoRealizada = true
							}
						}
					case 4:
						fmt.Println(conta.Extrato())
					case 5:
						fmt.Println(conta.PagarBoleto())
					case 6:
						os.Exit(0)
					default:
						fmt.Println("Opção inválida")
					}

				}
			}
		}
		if operacaoRealizada {
			continue
		}
	}
}

func DigiteConta() int {
	var cc int
	fmt.Println("Digite o número da conta: ")
	fmt.Scan(&cc)
	return cc
}
