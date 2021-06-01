package util

type LoginCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegistrationCredentials struct {
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password1 string `json:"password"`
	Password2 string `json:"confirm_password"`
}
