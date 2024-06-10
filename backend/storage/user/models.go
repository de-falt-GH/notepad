package user

import "time"

type DetailUserRequest struct {
	Id    int
	Login string
}

type DetailUserResponse struct {
	Login        string
	PasswordHash string
	Email        string
	Name         string
	Info         string
}

type UpdateUserRequest struct {
	Id           int
	Login        string
	PasswordHash string
	Email        string
	Name         string
	Info         string
}

type AddNoteRequest struct {
	UserId int
	Name   string
	Data   string
	Public bool
}

type AddNoteResponse struct {
	NoteId int
}

type UpdateNoteRequest struct {
	Id     int
	Name   string
	Data   string
	Public bool
}

type DetailNoteRequest struct {
	Id int
}

type DeleteNoteRequest struct {
	Id int
}

type ListPrivateNotesRequest struct {
	UserId int
	Search string
	Skip   int
	Limit  int
}

type Note struct {
	Id         int
	UserId     int
	Name       string
	Data       string
	Public     bool
	Created    time.Time
	Updated    time.Time
	AuthorName string
}
