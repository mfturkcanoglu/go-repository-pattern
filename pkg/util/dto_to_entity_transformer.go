package util

import (
	"github.com/mfturkcanoglu/go-repository-pattern/pkg/dto"
	"github.com/mfturkcanoglu/go-repository-pattern/pkg/entity"
)

func TodoDtoToEntity(dto dto.TodoDto) entity.Todo {
	return entity.Todo{
		Task:     dto.Task,
		TodoDate: dto.TodoDate,
	}
}
