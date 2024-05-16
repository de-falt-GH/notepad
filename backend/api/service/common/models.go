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
