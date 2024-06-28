package dto

import (
	"time"
)

type UserCreatedDTO struct {
	FirstName string `json:"firstname" form:"firstname" binding:"required"`
	LastName  string `json:"lastname" form:"lastname" binding:"required"`
	Email     string `json:"email" form:"email" binding:"required"`
	UserName  string `json:"username" form:"username" binding:"required"`
	Password  string `json:"password" form:"password" binding:"required"`
	RoleID    uint   `json:"role" form:"role" gorm:"default:0"`
}

type UserCreateAdminDTO struct {
	FirstName string `json:"firstname" form:"firstname" binding:"required"`
	LastName  string `json:"lastname" form:"lastname" binding:"required"`
	Email     string `json:"email" form:"email" binding:"required"`
	UserName  string `json:"username" form:"username" binding:"required"`
	Password  string `json:"password" form:"password" binding:"required"`
	RoleID    uint   `json:"role" form:"role" gorm:"default:0"`
}

type UserUpdateAdminDTO struct {
	ID        uint64 `json:"id" form:"id" binding:"required"`
	FirstName string `json:"first_name" form:"firstname" binding:"required"`
	LastName  string `json:"last_name" form:"lastname" binding:"required"`
	UserName  string `json:"username" form:"username" binding:"required"`
	Password  string `json:"password" form:"password"`
	Email     string `json:"email" form:"email" binding:"required"`
	RoleID    uint   `json:"role_id" form:"role" binding:"required"`
}

type UserResponseDTO struct {
	ID        uint64    `json:"id" form:"id" binding:"required"`
	UserName  string    `json:"username" form:"username" binding:"required"`
	Email     string    `json:"email" form:"email" binding:"required"`
	FirstName string    `json:"firstname" form:"firstname" binding:"required"`
	LastName  string    `json:"lastname" form:"lastname" binding:"required"`
	RoleName  string    `json:"role" form:"role" binding:"required"`
	CreatedAt time.Time `json:"createdAt" form:"createdAt" binding:"required"`
}

type UserUpdateDTO struct {
	UserID    string `json:"userid" form:"userid" binding:"required"`
	UserName  string `json:"username" form:"username"`
	Email     string `json:"email" form:"email"`
	FirstName string `json:"firstname" form:"firstname"`
	LastName  string `json:"lastname" form:"lastname"`
}

type UserUpdateResponseDTO struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}
