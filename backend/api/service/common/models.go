package common

type registerRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Info     string `json:"info"`
}

type authorizeRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type ListNotesRequest struct {
	Search string `form:"search"`
	Limit  int    `form:"limit"`
	Skip   int    `form:"skip"`
}

type NoteListItem struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type ListPublicNotesResponse struct {
	Notes []NoteListItem `json:"notes"`
}
