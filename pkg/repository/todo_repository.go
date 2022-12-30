package repository

import (
	"context"
	"log"

	"github.com/mfturkcanoglu/go-repository-pattern/pkg/database"
	"github.com/mfturkcanoglu/go-repository-pattern/pkg/entity"
)

type TodoRepository struct {
	logger *log.Logger
	client database.Client
}

func NewTodoRepository(l *log.Logger, c database.Client) Repository[entity.Todo] {
	return &TodoRepository{
		logger: l,
		client: c,
	}
}

// Find implements Repository
func (r *TodoRepository) Find(ctx context.Context, id string) (entity.Todo, error) {
	var todo entity.Todo
	res := r.client.GetDb().Model(&entity.Todo{}).First(&todo, id)
	return todo, res.Error
}

// FindAll implements Repository
func (r *TodoRepository) FindAll(context.Context) ([]entity.Todo, error) {
	var todos []entity.Todo
	res := r.client.GetDb().Model(&entity.Todo{}).Find(&todos)
	return todos, res.Error
}

// Remove implements Repository
func (r *TodoRepository) Remove(ctx context.Context, id string) error {
	res := r.client.GetDb().Delete(&entity.Todo{}, id)
	return res.Error
}

// Save implements Repository
func (r *TodoRepository) Save(ctx context.Context, todo entity.Todo) (entity.Todo, error) {
	res := r.client.GetDb().Save(&todo)
	return todo, res.Error
}

// Update implements Repository
func (r *TodoRepository) Update(ctx context.Context, id string, todo entity.Todo) (entity.Todo, error) {
	oldTodo, err := r.Find(ctx, id)
	if err != nil {
		return oldTodo, err
	}
	res := r.client.GetDb().Model(&entity.Todo{}).Updates(todo)
	return todo, res.Error
}
