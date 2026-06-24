package services

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"time"

	"github.com/yhlas/basic-pharmacy/internal/models"
	"github.com/yhlas/basic-pharmacy/internal/repositories"
)

func UserListService(c context.Context, filter repositories.UserFilter) ([]models.UserResponse, error) {

	list, err := repositories.UserList(c, filter)
	if err != nil {
		return nil, err
	}
	res := []models.UserResponse{}
	for _, v := range list {
		item := models.UserResponse{}
		item.ID = v.ID
		item.Name = v.Name
		item.Role = v.Role
		item.Password = v.Password
		item.Email = v.Email
		// ...
		res = append(res, item)
	}
	return res, nil
}
func DeleteUserService(c context.Context, id int) error {
	return repositories.UserDelete(c, id)
}
func GetUserService(c context.Context, token string, word bool) (models.UserResponse, error) {
	return repositories.GetUser(c, token, word)
}
func UpdateUserService(c context.Context, token string, req models.UserUpdateRequest) error {
	return repositories.UpdateUser(c, token, req)
}
func GenerateSecureToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}

func RegistrationService(c context.Context, name string, role string, password string, email string, phone string, region string) error {
	return repositories.InsertUser(c, name, phone, region, email, password)
}

func LoginService(c context.Context, email string, password string) (string, error) {
	user, err := repositories.UserGetByEmail(c, email)
	if err != nil {
		return "", err
	}
	if user.Password != password {
		return "", errors.New("wrong password")
	}
	Token := GenerateSecureToken(32)

	expiry := time.Now().Add(30 * 6 * 24 * time.Hour)

	err = repositories.InsertToken(c, user.ID, Token, expiry)
	if err != nil {
		return "", err
	}
	return Token, nil
}
func LogoutService(c context.Context, token string) (string, error) {
	if repositories.CheckIsTokenReal(c, token) {
		err := repositories.DeleteToken(c, token)
		if err != nil {
			return "", err
		}
		return "logout", nil
	}
	return "not token", nil
}
func ChangePasswordService(c context.Context, token string, word bool, passchange models.ChangePasswordRequest, req models.UserResponse) error {
	req, err := repositories.GetUser(c, token, word)
	if err != nil {
		return err
	}
	if req.Password != passchange.OldPassword {
		return errors.New("wrong pass")
	}
	err = repositories.UpdatePassword(c, token, passchange)
	if err != nil {
		return err
	}
	return nil
}
