package handlers

import (
	"net/http"
	"note/internal/models"
	"note/internal/service"
	utils "note/pkg"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	todoService *service.TodoService
}

func NewTodoHandler(todoService *service.TodoService) *TodoHandler {
	return &TodoHandler{
		todoService: todoService,
	}
}

func (h *TodoHandler) GetAllTodos(c *gin.Context) {
	user := c.MustGet("user").(*models.User)
	todos, err := h.todoService.GetAllTodos(user)
	if err != nil {
		utils.ErrorJSON(c, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(c, http.StatusOK, todos)
}

func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var req struct {
		Title  string `json:"title"`
		Detail string `json:"detail"`
	}
	if err := utils.ReadJSON(c, &req); err != nil {
		utils.ErrorJSON(c, err, http.StatusBadRequest)
		return
	}
	user := c.MustGet("user").(*models.User)
	todo, err := h.todoService.CreateTodo(user, req.Title, req.Detail)
	if err != nil {
		utils.ErrorJSON(c, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(c, http.StatusCreated, todo)
}

func (h *TodoHandler) FindTodo(c *gin.Context) {
	user := c.MustGet("user").(*models.User)
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorJSON(c, err, http.StatusBadRequest)
		return
	}
	todo, err := h.todoService.FindTodo(user, id)
	if err != nil {
		utils.ErrorJSON(c, err, http.StatusNotFound)
		return
	}
	utils.WriteJSON(c, http.StatusOK, todo)
}

func (h *TodoHandler) MarkTodoComplete(c *gin.Context) {
	user := c.MustGet("user").(*models.User)
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorJSON(c, err, http.StatusBadRequest)
		return
	}
	todo, err := h.todoService.FindTodo(user, id)
	if err != nil {
		utils.ErrorJSON(c, err, http.StatusNotFound)
		return
	}

	var req struct {
		Complete bool `json:"complete"`
	}
	if err := utils.ReadJSON(c, &req); err != nil {
		utils.ErrorJSON(c, err, http.StatusBadRequest)
		return
	}

	err = h.todoService.MarkTodoComplete(todo, req.Complete)
	if err != nil {
		utils.ErrorJSON(c, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(c, http.StatusOK, todo)
}
