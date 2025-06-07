package main

import (
	"fmt"
	"os"
	"taskmanager/entity"
	"taskmanager/services"
	"taskmanager/utils"
)

func main() {

	var minhasTarefas services.TaskService

	for {
		fmt.Println("1 - Criar tarefa")
		fmt.Println("2 - Listar tarefas")
		fmt.Println("3 - Concluir tarefa")
		fmt.Println("4 - Remover tarefa")
		fmt.Println("5 - Sair \n")
		fmt.Print("Escolha uma opção: ")

		var opcao string
		fmt.Scanln(&opcao)

		switch opcao {

		case "1":

			var Title string
			var Completed bool

			fmt.Println("Qual o titulo da Tarefa: ")
			fmt.Scanln(&Title)

			id := utils.GerarIdAleatorio()

		case "2":

		case "3":

		case "4":

		case "5":
			fmt.Println("Saindo do programa..!")
			os.Exit(5)

		default:
			fmt.Println("Opção inválida, tente novamente!")

		}
	}
}
