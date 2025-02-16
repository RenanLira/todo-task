package services

import (
	"todo-tasks/internal/domain/todos"
	"todo-tasks/internal/utils/types"
)



type GetAllTodosResponseDTO struct {
	Todos []*todos.Todo `json:"todos"`
	Page  types.Page          `json:"pageInfo"`
}

func (s *TodoService) GetAllTodos(dto todos.ReqGetAllTodosDTO) (*GetAllTodosResponseDTO, error) {

	todos, err := s.TodoRepository.GetAll(dto.Limit, dto.Page*dto.Limit)
	if err != nil {
		return nil, err
	}

	pageInfo, err := s.TodoRepository.GetPageInfo(int(dto.Limit), int(dto.Page))
	if err != nil {
		return nil, err
	}

	return &GetAllTodosResponseDTO{
		Todos: todos,
		Page:  pageInfo,
	}, nil
}
