package main

import "fmt"

func main() {

	var opcao int

	fmt.Println("1 - Criar tarefa")
	fmt.Println("2 - Listar tarefas")
	fmt.Println("3 - Concluir tarefa")
	fmt.Println("4 - Remover tarefa")
	fmt.Println("5 - Sair")
	fmt.Print("Escolha uma opção: ")
	fmt.Scan(&opcao)

	switch opcao {

	case "1":
		fmt.Println()

	}

}
