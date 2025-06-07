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

			fmt.Print("Qual o titulo da Tarefa: ")
			fmt.Scanln(&Title)

			id := utils.GerarIdAleatorio()
			fmt.Println("O seu ID aleatório: ", id)

			fmt.Println("Tarefa criada com sucesso! \n")

			novaTarefa := entity.Task{
				ID:       id,
				Titulo:   Title,
				Completo: false,
			}

			minhasTarefas.CriarTarefa(novaTarefa)

		case "2":
			for _, tarefa := range minhasTarefas.ListarTarefas() {
				fmt.Printf("ID: %s, Titulo: %s, Completo: %t \n", tarefa.ID, tarefa.Titulo, tarefa.Completo)
				fmt.Println("")
			}

		case "3":

			var id string

			fmt.Print("Qual ID da Tarefa para atualização de Status: ")
			fmt.Scanln(&id)

			valido := minhasTarefas.TarefaConcluida(id)

			if valido == true {
				fmt.Println("Tarefa Concluida! \n")
			} else {
				fmt.Println("Id não encontrado, tente novamente.")
			}

		case "4":

			var id string

			fmt.Print("Qual ID da Tarefa que você quer remover: ")
			fmt.Scanln(&id)

			mensagem := minhasTarefas.RemoverTarefaPeloId(id)
			fmt.Println(mensagem)

		case "5":
			fmt.Println("Saindo do programa..!")
			os.Exit(5)

		default:
			fmt.Println("Opção inválida, tente novamente!")

		}
	}
}
