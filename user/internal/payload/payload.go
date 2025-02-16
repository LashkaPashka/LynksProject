package payload

type RegisterUserRequest struct{
	Email string `json:"email"`
	Name string `json:"name"`
	Password string `json:"password"`
}

type LoginUserRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}