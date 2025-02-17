package todos

type CreateTodoDTO struct {
	Text   string `json:"text"`
	UserID string `json:"user_id"`
}

type ReqGetAllTodosDTO struct {
	Limit  int32  `json:"limit"`
	Page   int32  `json:"offset"`
	Search string `json:"search"`
}
