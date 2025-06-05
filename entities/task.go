package entities

type Task struct {
	ID        string
	Title     string
	Completed bool
}
type TaskService struct {
	Tarefas []Task
}
