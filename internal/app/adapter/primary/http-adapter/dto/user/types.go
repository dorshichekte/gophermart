package dto

type AuthRequest struct {
	Login    string `json:"login" validate:"required,min=8,max=255"`
	Password string `json:"password" validate:"required,min=8,max=255"`
}

type RegisterRequest struct {
	AuthRequest
}

type LoginRequest struct {
	AuthRequest
}
