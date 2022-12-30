package service

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/mfturkcanoglu/go-repository-pattern/pkg/dto"
	"github.com/mfturkcanoglu/go-repository-pattern/pkg/entity"
	"github.com/mfturkcanoglu/go-repository-pattern/pkg/repository"
	"github.com/mfturkcanoglu/go-repository-pattern/pkg/util"
)

type TodoService struct {
	logger     *log.Logger
	repository repository.Repository[entity.Todo]
}

func NewTodoService(l *log.Logger, r repository.Repository[entity.Todo]) Service[entity.Todo, dto.TodoDto] {
	return &TodoService{
		logger:     l,
		repository: r,
	}
}

// Create implements Service
func (s *TodoService) Create(ctx context.Context, dto dto.TodoDto) (entity.Todo, error) {
	todo := util.TodoDtoToEntity(dto)
	todo.ID = uuid.New().String()
	todo.CreateDate = time.Now()
	return s.repository.Save(ctx, todo)
}

// Delete implements Service
func (s *TodoService) Delete(ctx context.Context, id string) error {
	return s.repository.Remove(ctx, id)
}

// Get implements Service
func (s *TodoService) Get(ctx context.Context, id string) (entity.Todo, error) {
	return s.repository.Find(ctx, id)
}

// GetAll implements Service
func (s *TodoService) GetAll(ctx context.Context) ([]entity.Todo, error) {
	return s.repository.FindAll(ctx)
}

// Update implements Service
func (s *TodoService) Update(ctx context.Context, id string, todoDto dto.TodoDto) (entity.Todo, error) {
	todo := util.TodoDtoToEntity(todoDto)
	return s.repository.Update(ctx, id, todo)
}
