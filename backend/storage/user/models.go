package user

type DetailUserRequest struct {
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

type ListNotesRequest struct {
	UserId int
	Skip   int
	Limit  int
}

type Note struct {
	Id     int
	Name   string
	Data   string
	Public bool
}
