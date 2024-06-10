package common

import "time"

type CreateUserRequest struct {
	Login        string
	PasswordHash string
	Email        string
	Name         string
	Info         string
}

type DetailUserRequest struct {
	Login string
}

type DetailUserResponse struct {
	Id           int
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

type ListPublicNotesRequest struct {
	Search string
	Skip   int
	Limit  int
}

type DetailNoteRequest struct {
	Id int
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
