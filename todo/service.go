package todo

type todoService struct {
	todoRepository TodoRepository
}

func (s todoService) createTodo(label string) (Todo, error) {
	todo := Todo{Label: label, IsComplete: false}

	id, err := s.todoRepository.Insert(todo.Label, todo.IsComplete)
	if err != nil {
		return todo, err
	}

	// Set the id
	todo.ID = id

	return todo, nil
}

func (s todoService) getTodoByID(id interface{}) (Todo, error) {
	todo := Todo{ID: id}
	err := s.todoRepository.GetByID(id, &todo.Label, &todo.IsComplete)
	return todo, err
}
