package services

import (
	"sync"
	"todo-tasks/internal/domain/todos"
	"todo-tasks/internal/utils/types"
)

type GetAllTodosResponseDTO struct {
	Todos []*todos.Todo `json:"todos"`
	Page  types.Page    `json:"pageInfo"`
}

func (t *TodoService) GetAllTodos(dto todos.ReqGetAllTodosDTO, userId string) (*GetAllTodosResponseDTO, error) {

	var wg sync.WaitGroup

	var all []*todos.Todo
	var pageInfo types.Page
	var err error

	wg.Add(1)
	go func() {
		defer wg.Done()
		all, err = t.TodoRepository.GetAllByUser(userId, dto.Limit, dto.Page*dto.Limit)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		pageInfo, err = t.TodoRepository.GetPageInfo(int(dto.Limit), int(dto.Page))
	}()

	wg.Wait()

	if err != nil {
		return nil, err
	}

	return &GetAllTodosResponseDTO{
		Todos: all,
		Page:  pageInfo,
	}, nil
}
