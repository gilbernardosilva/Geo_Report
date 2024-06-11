package dto

type UserCreatedDTO struct {
	FirstName string `json:"firstname" form:"firstname" binding:"required"`
	LastName  string `json:"lastname" form:"lastname" binding:"required"`
	Email     string `json:"email" form:"email" binding:"required"`
	UserName  string `json:"username" form:"username" binding:"required"`
	Password  string `json:"password" form:"password" binding:"required"`
}

type UserResponseDTO struct {
	UserName  string `json:"username" form:"username" binding:"required"`
	Email     string `json:"email" form:"email" binding:"required"`
	FirstName string `json:"firstname" form:"firstname" binding:"required"`
	LastName  string `json:"lastname" form:"lastname" binding:"required"`
}
