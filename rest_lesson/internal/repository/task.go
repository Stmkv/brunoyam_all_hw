package repository

import "rest_lesson/internal/model"

type TaskRepository interface {
	GetAll() []model.Task
	Create(task model.Task) model.Task
	GetById(id int) (model.Task, error)
}
