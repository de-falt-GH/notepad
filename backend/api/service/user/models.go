package user

type DetailUserRequest struct {
	Id int `json:"id"`
}

type DetailUserResponse struct {
	Login string `json:"login"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Info  string `json:"info"`
}

type UpdateUserRequest struct {
	Id       string `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Info     string `json:"info"`
}

type AddNoteRequest struct {
	Name   string `json:"name"`
	Data   string `json:"data"`
	Public bool   `json:"public"`
}

type UpdateNoteRequest struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Data   string `json:"data"`
	Public bool   `json:"public"`
}

type DetailNoteRequest struct {
	Id int `json:"id"`
}

type DeleteNoteRequest struct {
	Id int `json:"id"`
}

type ListNotesRequest struct {
	limit  int  `json:"limit"`
	skip   int  `json:"skip"`
	public bool `json:"public"`
}
