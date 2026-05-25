package repositories

import (
	"context"

	"github.com/yhlas/basic-pharmacy/internal/models"
	"github.com/yhlas/basic-pharmacy/internal/utils"
)

func Registration(c context.Context, user models.User) (models.User, error) {

	_, err := utils.GetDB().Exec(context.Background(),
		"INSERT INTO users(id, name, email, password, role) VALUES ($1,$2,$3,$4,$5)",
		user.ID, user.Name, user.Email, user.Password, user.Role,
	)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
