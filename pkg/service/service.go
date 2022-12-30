package service

import "context"

type Service[T any, TDto any] interface {
	GetAll(context.Context) ([]T, error)
	Get(context.Context, string) (T, error)
	Create(context.Context, TDto) (T, error)
	Update(context.Context, string, TDto) (T, error)
	Delete(context.Context, string) error
}
