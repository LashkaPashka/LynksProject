package payload

type UserRequest struct{
	Email string `json:"email"`
	Name string `json:"name"`
	Password string `json:"password"`
}