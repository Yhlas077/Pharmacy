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

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type UserResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}
