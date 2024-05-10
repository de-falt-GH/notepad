package common

type registerRequest struct {
	login    string
	password string
	email    string
	name     string
	info     string
}

type authorizeRequest struct {
	login    string
	password string
}
