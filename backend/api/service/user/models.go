package user

type DetailUserRequest struct {
	id int
}

type DetailUserResponse struct {
	login string
	email string
	name  string
	info  string
}

type UpdateUserRequest struct {
	id       string
	login    string
	password string
	email    string
	name     string
	info     string
}

type AddNoteRequest struct {
	name   string
	data   string
	public bool
}

type UpdateNoteRequest struct {
	id     int
	name   string
	data   string
	public bool
}

type DetailNoteRequest struct {
	id int
}

type DeleteNoteRequest struct {
	id string
}

type ListNotesRequest struct {
	limit    int
	skip     int
	personal bool
}
