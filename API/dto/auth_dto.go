package dto

type LoginDTO struct {
	UserName string `json:"username" form:"username" binding:"required" validate:"email"`
	Password string `json:"password" form:"password,omitempty" validate:"min:6" binding:"required"`
}

type LoginResponse struct {
	Token   string `json:"token"`
	Message string `json:"message"`
}
