package main

import (
	"rest_lesson/internal/handler"
	"rest_lesson/internal/repository"
	"rest_lesson/internal/service"

	"github.com/gin-gonic/gin"
)

func getTaskHandler() handler.TaskHandler {
	taskRepo := repository.NewInMemoryTaskRepository()
	taskService := service.NewTaskService(taskRepo)
	taskHandler := handler.NewTaskHandler(taskService)
	return *taskHandler
}

func main() {
	r := gin.Default()
	taskHandler := getTaskHandler()

	r.GET("/hello", handler.HelloHandler)
	r.GET("/tasks", taskHandler.GetTasksHandler)
	r.POST("/task", taskHandler.CreateTaskHandler)

	err := r.Run(":8080")
	if err != nil {
		// !TODO Посмотреть как ошибку в stdout выкинуть вместе с текстом
		panic(`gin run error`)
	}
}
