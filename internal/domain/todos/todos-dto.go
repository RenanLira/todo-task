package todos

type CreateTodoDTO struct {
	Text string `json:"text"`
}

type ReqGetAllTodosDTO struct {
	Limit  int32  `json:"limit"`
	Page   int32  `json:"offset"`
	Search string `json:"search"`
}

