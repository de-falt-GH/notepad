package common

import "time"

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
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	AuthorName string    `json:"author_name"`
	Updated    time.Time `json:"updated"`
}

type ListPublicNotesResponse struct {
	Notes []NoteListItem `json:"notes"`
}
