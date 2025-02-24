package services

import (
	"todo-tasks/internal/domain/todos"
	"todo-tasks/internal/utils/types"
)

type GetAllTodosResponseDTO struct {
	Todos []*todos.Todo `json:"todos"`
	Page  types.Page    `json:"pageInfo"`
}

func (t *TodoService) GetAllTodos(dto todos.ReqGetAllTodosDTO, userId string) (*GetAllTodosResponseDTO, error) {

	all, err := t.TodoRepository.GetAllByUser(userId, dto.Limit, dto.Page*dto.Limit)
	if err != nil {
		return nil, err
	}

	pageInfo, err := t.TodoRepository.GetPageInfo(int(dto.Limit), int(dto.Page))
	if err != nil {
		return nil, err
	}

	return &GetAllTodosResponseDTO{
		Todos: all,
		Page:  pageInfo,
	}, nil
}
