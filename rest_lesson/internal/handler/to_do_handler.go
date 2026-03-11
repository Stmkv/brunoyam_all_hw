package handler

import (
	"net/http"
	"rest_lesson/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	service *service.TaskService
}

func NewTaskHandler(s *service.TaskService) *TaskHandler {
	return &TaskHandler{service: s}
}

func (h *TaskHandler) GetTasksHandler(c *gin.Context) {
	tasks := h.service.GetTasks()

	c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) CreateTaskHandler(c *gin.Context) {
	var req struct {
		Title string `json:"title"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	task := h.service.CreateTask(req.Title)

	c.JSON(http.StatusCreated, task)
}

func (h *TaskHandler) GetById(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)   // TODO обработать ошибку
	task, _ := h.service.GetById(id) // TODO обработать ошибку
	c.JSON(http.StatusOK, task)
}
