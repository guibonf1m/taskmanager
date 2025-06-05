package services

import (
	"fmt"
	"taskmanager/entity"
)

type TaskService struct {
	Tarefas []entity.Task // / Armazena todas as tarefas criadas; cada tarefa utiliza a
	// estrutura Task definida no pacote entity para manter o código organizado,
	//reaproveitável e separado por responsabilidade.
}

func (t *TaskService) CriarTarefa(task entity.Task) {
	t.Tarefas = append(t.Tarefas, task)
}

func (t *TaskService) ListarTarefas() []entity.Task {
	return t.Tarefas
}

func (t *TaskService) TarefaConcluida(idDesejado string) bool {
	// Percorre todas as tarefas pelo índice e cópia da tarefa.
	for i, tarefa := range t.Tarefas {
		// Compara o ID da tarefa atual com o ID buscado.
		if tarefa.ID == idDesejado {
			// Marca a tarefa como concluída alterando o campo Completed para true.
			t.Tarefas[i].Completed = true
			// Retorna true para indicar que a tarefa foi encontrada e alterada.
			return true
		}
	}
	// Retorna false caso nenhuma tarefa com o ID tenha sido encontrada.
	return false
}

func (t *TaskService) RemoverTarefaPeloId(idRecebido string) string {
	for i, id := range t.Tarefas {
		if id.ID == idRecebido {
			t.Tarefas = append(t.Tarefas[:i], t.Tarefas[i+1:]...)
			return fmt.Sprintf("Id encontrado e removido %s", idRecebido)
		}
	}
	return fmt.Sprint("Id não encontrado, tente novamente")
}
