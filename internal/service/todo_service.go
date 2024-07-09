package service

import (
	"note/internal/models"
	"note/internal/repository"
)

type TodoService struct {
	todoRepo repository.TodoRepository
}

func NewTodoService(todoRepo repository.TodoRepository) *TodoService {
	return &TodoService{
		todoRepo: todoRepo,
	}
}

func (s *TodoService) GetAllTodos(user *models.User) ([]models.Todo, error) {
	return s.todoRepo.GetAll(user.ID)
}

func (s *TodoService) CreateTodo(user *models.User, title, detail string) (*models.Todo, error) {
	todo := &models.Todo{Title: title, Detail: detail, UserID: user.ID, Complete: false}
	err := s.todoRepo.Create(todo)
	return todo, err
}

func (s *TodoService) FindTodo(user *models.User, id uint64) (*models.Todo, error) {
	return s.todoRepo.Find(user.ID, id)
}

func (s *TodoService) MarkTodoComplete(todo *models.Todo, state bool) error {
	todo.Complete = state
	return s.todoRepo.Update(todo)
}
