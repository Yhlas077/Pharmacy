package models

type Token struct {
	ID    int
	Token string
}
type TokenCreateRequest struct {
	Password string
	Email    string
}
type TokenResponse struct {
	Token  string
	UserID int
}
type TokenCheck struct {
	Token
}
