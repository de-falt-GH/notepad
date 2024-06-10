package user

import "time"

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
	Name   string `json:"name"`
	Data   string `json:"data"`
	Public bool   `json:"public"`
}

type ListNotesRequest struct {
	Search string `form:"search"`
	Limit  int    `form:"limit"`
	Skip   int    `form:"skip"`
}

type NoteListItem struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	AuthorName string    `json:"author_name"`
	Updated    time.Time `json:"updated"`
}

type ListPrivateNotesResponse struct {
	Notes []NoteListItem `json:"notes"`
}
