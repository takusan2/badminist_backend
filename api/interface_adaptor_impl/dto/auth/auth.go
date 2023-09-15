package auth_dto

type TemporaryRegistrationRequestBody struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ActivateUserRequestBody struct {
	Email       string `json:"email"`
	ConfirmPass string `json:"confirm_pass"`
}

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponseBody struct {
	Token string `json:"token"`
}
