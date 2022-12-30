package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mfturkcanoglu/go-repository-pattern/pkg/dto"
	"github.com/mfturkcanoglu/go-repository-pattern/pkg/entity"
	"github.com/mfturkcanoglu/go-repository-pattern/pkg/service"
)

type TodoController struct {
	logger  *log.Logger
	service service.Service[entity.Todo, dto.TodoDto]
}

func NewTodoController(l *log.Logger, s service.Service[entity.Todo, dto.TodoDto]) Controller {
	return &TodoController{
		logger:  l,
		service: s,
	}
}

// Delete implements Controller
func (c *TodoController) Delete(rw http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := c.service.Delete(r.Context(), id); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(dto.ErrorResponse{
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
		})
		return
	}
	json.NewEncoder(rw).Encode(dto.RestResponse{
		Success:    true,
		Data:       true,
		StatusCode: http.StatusOK,
	})
}

// GetByID implements Controller
func (c *TodoController) GetByID(rw http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	todo, err := c.service.Get(r.Context(), id)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(dto.ErrorResponse{
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
		})
		return
	}
	json.NewEncoder(rw).Encode(&dto.RestResponse{
		Success:    true,
		Data:       todo,
		StatusCode: http.StatusOK,
	})
}

// ListAll implements Controller
func (c *TodoController) ListAll(rw http.ResponseWriter, r *http.Request) {
	todos, err := c.service.GetAll(r.Context())
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(&dto.ErrorResponse{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		})
		return
	}
	json.NewEncoder(rw).Encode(dto.RestResponse{
		Success:    true,
		Data:       todos,
		StatusCode: http.StatusOK,
	})
}

// Post implements Controller
func (c *TodoController) Post(rw http.ResponseWriter, r *http.Request) {
	var todoDto dto.TodoDto
	if err := json.NewDecoder(r.Body).Decode(&todoDto); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(dto.ErrorResponse{
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
		})
		return
	}
	todo, err := c.service.Create(r.Context(), todoDto)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(dto.ErrorResponse{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		})
		return
	}
	json.NewEncoder(rw).Encode(dto.RestResponse{
		Success:    true,
		Data:       todo,
		StatusCode: http.StatusCreated,
	})
}

// Put implements Controller
func (c *TodoController) Put(rw http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var todoDto dto.TodoDto
	if err := json.NewDecoder(r.Body).Decode(&todoDto); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(dto.ErrorResponse{
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
		})
		return
	}
	todo, err := c.service.Update(r.Context(), id, todoDto)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(dto.ErrorResponse{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		})
		return
	}
	json.NewEncoder(rw).Encode(dto.RestResponse{
		Success:    true,
		Data:       todo,
		StatusCode: http.StatusOK,
	})
}
