package common

type registerRequest struct {
	Login    string
	Password string
	Email    string
	Name     string
	Info     string
}

type authorizeRequest struct {
	Login    string
	Password string
}
