package repository

import (
	"errors"
	"rest_lesson/internal/model"
)

type InMemoryTaskRepository struct {
	tasks []model.Task
}

func NewInMemoryTaskRepository() *InMemoryTaskRepository {
	return &InMemoryTaskRepository{
		tasks: []model.Task{
			{ID: 1, Title: "first", Done: false},
			{ID: 2, Title: "second", Done: false},
		},
	}
}

func (r *InMemoryTaskRepository) GetAll() []model.Task {
	return r.tasks
}

func (r *InMemoryTaskRepository) Create(task model.Task) model.Task {
	task.ID = len(r.tasks) + 1
	r.tasks = append(r.tasks, task)
	return task
}

func (r *InMemoryTaskRepository) GetById(id int) (model.Task, error) {
	for _, task := range r.tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return model.Task{}, errors.New("task not found")
}
