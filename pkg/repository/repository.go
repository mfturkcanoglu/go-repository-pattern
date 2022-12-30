package repository

import "context"

type Repository[T any] interface {
	FindAll(context.Context) ([]T, error)
	Find(context.Context, string) (T, error)
	Save(context.Context, T) (T, error)
	Update(context.Context, string, T) (T, error)
	Remove(context.Context, string) error
}
