package main

import (
	"Bank/clientes"
	"Bank/contas"
	"fmt"
	"os"
)

func menu() {
	fmt.Println("\nEscolha a operação que deseja realizar:")
	fmt.Println("1 - Criar Conta Corrente")
	fmt.Println("2 - Criar Conta Poupança")
	fmt.Println("3 - Sacar")
	fmt.Println("4 - Depositar")
	fmt.Println("5 - Transferir")
	fmt.Println("6 - Obter Saldo")
	fmt.Println("0 - Sair")
}

func main() {
	var comando int
	var contasCorrentes = make(map[string]*contas.ContaCorrente) // Mapa para armazenar contas correntes
	var contasPoupanca = make(map[string]*contas.ContaPoupanca)  // Mapa para armazenar contas poupança

	for {
		menu()
		fmt.Print("Digite o comando que deseja: ")
		fmt.Scanln(&comando)
		fmt.Println(" ")

		switch comando {
		case 1:
			var nome, cpf, profissao string
			var numAgencia, numConta int
			fmt.Print("Digite o nome da sua conta: ")
			fmt.Scanln(&nome)
			fmt.Print("Digite o CPF da sua conta: ")
			fmt.Scanln(&cpf)
			fmt.Print("Digite a profissão da sua conta: ")
			fmt.Scanln(&profissao)
			fmt.Print("Digite o número da agência da sua conta: ")
			fmt.Scanln(&numAgencia)
			fmt.Print("Digite o número da conta da sua conta: ")
			fmt.Scanln(&numConta)

			// Criação da conta corrente
			conta := &contas.ContaCorrente{
				Titular: clientes.Titular{
					Nome:      nome,
					Cpf:       cpf,
					Profissao: profissao,
				},
				NumAgencia:  numAgencia,
				NumeroConta: numConta,
			}
			contasCorrentes[cpf] = conta // Adiciona a conta ao mapa usando o CPF como chave
			fmt.Println("Conta corrente criada com sucesso!")

		case 2:
			var nome, cpf, profissao string
			var numAgencia, numConta int
			fmt.Print("Digite o nome da sua conta: ")
			fmt.Scanln(&nome)
			fmt.Print("Digite o CPF da sua conta: ")
			fmt.Scanln(&cpf)
			fmt.Print("Digite a profissão da sua conta: ")
			fmt.Scanln(&profissao)
			fmt.Print("Digite o número da agência da sua conta: ")
			fmt.Scanln(&numAgencia)
			fmt.Print("Digite o número da conta da sua conta: ")
			fmt.Scanln(&numConta)

			// Criação da conta corrente
			conta := &contas.ContaPoupanca{
				Titular: clientes.Titular{
					Nome:      nome,
					Cpf:       cpf,
					Profissao: profissao,
				},
				NumAgencia:  numAgencia,
				NumeroConta: numConta,
			}
			contasPoupanca[cpf] = conta // Adiciona a conta ao mapa usando o CPF como chave
			fmt.Println("Conta poupança criada com sucesso!")

		case 3:
			var cpf string
			var valor float64
			fmt.Print("Digite o CPF da conta: ")
			fmt.Scanln(&cpf)
			fmt.Print("Digite o valor a sacar: ")
			fmt.Scanln(&valor)

			if conta, exists := contasCorrentes[cpf]; exists {
				result := conta.Sacar(valor)
				fmt.Println(result)
			} else {
				fmt.Println("Conta não encontrada!")
			}

		case 4:
			var cpf string
			var valor float64
			fmt.Print("Digite o CPF da conta: ")
			fmt.Scanln(&cpf)
			fmt.Print("Digite o valor a depositar: ")
			fmt.Scanln(&valor)

			if conta, exists := contasCorrentes[cpf]; exists {
				conta.Depositar(valor)
				fmt.Println("Depósito realizado com sucesso!")
			} else {
				fmt.Println("Conta não encontrada!")
			}

		case 5:
			var valor float64
			var cpf1, cpf2 string
			fmt.Println("Digite o CPF da sua conta corrente: ")
			fmt.Scanln(&cpf1)
			fmt.Println("Digite o CPF da conta corrente de destino: ")
			fmt.Scanln(&cpf2)
			fmt.Print("Digite o valor a transferir: ")
			fmt.Scanln(&valor)

			// Verifica se as contas existem no mapa
			contaOrigem, exists1 := contasCorrentes[cpf1]
			contaDestino, exists2 := contasCorrentes[cpf2]

			// Confirma se ambas as contas existem
			if exists1 && exists2 {
				// Verifica se há saldo suficiente na conta de origem
				if contaOrigem.ObterSaldo() >= valor {
					contaOrigem.Sacar(valor)
					contaDestino.Depositar(valor)
					fmt.Println("Transferência realizada com sucesso!")
				} else {
					fmt.Println("Saldo insuficiente na conta de origem!")
				}
			} else {
				// Caso uma das contas não exista, exibe a mensagem de erro
				if !exists1 {
					fmt.Println("Conta de origem não encontrada!")
				}
				if !exists2 {
					fmt.Println("Conta de destino não encontrada!")
				}
			}

		case 6:
			var cpf string
			fmt.Print("Digite o CPF da conta: ")
			fmt.Scanln(&cpf)

			if conta, exists := contasCorrentes[cpf]; exists {
				fmt.Printf("Saldo: %.2f\n", conta.ObterSaldo())
			} else {
				fmt.Println("Conta não encontrada!")
			}

		case 0:
			fmt.Println("Saindo...")
			os.Exit(0)

		default:
			fmt.Println("Comando não identificado.")
		}
	}
}
