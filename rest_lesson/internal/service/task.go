package service

import (
	"rest_lesson/internal/model"

	"rest_lesson/internal/repository"
)

type TaskService struct {
	repo repository.TaskRepository
}

func NewTaskService(r repository.TaskRepository) *TaskService {
	return &TaskService{
		repo: r,
	}
}

func (s *TaskService) GetTasks() []model.Task {
	return s.repo.GetAll()
}

func (s *TaskService) CreateTask(title string) model.Task {
	task := model.Task{
		Title: title,
		Done:  false,
	}

	return s.repo.Create(task)
}

func (s *TaskService) GetById(id int) (model.Task, error) {
	return s.repo.GetById(id)
}
