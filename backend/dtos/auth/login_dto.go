package authdto

type LoginDto struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginDtoResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
