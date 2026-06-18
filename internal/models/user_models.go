package models

type Role string

const AdminRole Role = "admin"
const PharmacyRole Role = "pharmacy"
const UserRole Role = "user"

type ErrorResponse struct {
	Success   bool   `json:"success"`
	ErrorMsg  string `json:"error_msg" `
	ErrorCode string `json:"error_code"`
}

type UserCreateRequest struct {
	Name     string `json:"name"`
	Role     string `json:"role"`
	Password string `json:"password" binding:"required,min=8"`
	Email    string `json:"email" binding:"required,email"`
}

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Region   string `json:"region"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"oldpassword" binding:"required,min=8"`
	NewPassword string `json:"newpassword" binding:"required,min=8"`
}

type UserUpdateRequest struct {
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	Region string `josn:"region"`
}

type Meta struct {
	Total int `json:"total"`
	Limit int `json:"limit"`
	Offset int `json:"offset"`
}