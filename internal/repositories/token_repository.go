package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/yhlas/basic-pharmacy/internal/models"
)

//	func InsertToken(c context.Context, id int, token string) error {
//		_, err := GetDB().Exec(c, "INSERT into tokens(user_id, token) values ($1, $2)", id, token)
//		if err != nil {
//			return err
//		}
//		return nil
//	}
func InsertToken(c context.Context, id int, token string, expiresAt time.Time) error {
	// Insert the values directly, including the passed-in expiration time
	query := "INSERT INTO tokens (user_id, token, expires_at) VALUES ($1, $2, $3)"
	_, err := GetDB().Exec(c, query, id, token, expiresAt)
	if err != nil {
		return err
	}
	return nil
}

func GetToken(c context.Context, token string, id int) {
	row := GetDB().QueryRow(c, `select token from tokens where id =$1`, id)
	row.Scan(&token)
}

func DeleteToken(c context.Context, token string) error {
	var id int
	row := GetDB().QueryRow(c, `select user_id from tokens where token = $1`, token)
	row.Scan(&id)
	_, err := GetDB().Exec(c, `DELETE FROM tokens WHERE user_id=$1`, id)
	if err != nil {
		return err
	}
	return nil
}

func InsertUser(c context.Context, name string, phone string, region string, email string, password string) error {
	_, err := GetDB().Exec(c,
		"INSERT INTO users(name, phone, region, email, password, role) VALUES ($1,$2,$3,$4,$5,$6)",
		name, phone, region, email, password, "user",
	)

	if err != nil {
		fmt.Println("DB ERROR:", err)
		return err
	}

	return nil
}

type TokenResponse struct {
	Token  string
	UserID int
}

type Token struct {
	ID    int
	Token string
}

type TokenCheck struct {
	Token
}

func CheckIsTokenReal(c context.Context, token string) bool {
	db := GetDB()
	var Token TokenCheck
	rows, err := db.Query(c, "select token from tokens where token=$1", token)
	if err != nil {
		return false
	}
	err = rows.Scan(&Token.Token)
	if err != nil {
		return true
	}
	return false
}

func GetUserIdByToken(c context.Context, token string) (int, error) {
	db := GetDB()
	var req models.TokenResponse
	rows := db.QueryRow(c, "select token, user_id from tokens where token=$1", token)
	err := rows.Scan(&req.Token, &req.UserID)
	if err != nil {
		return 0, err
	}
	return req.UserID, nil
}

func GetExpiresAtByToken(c context.Context, token string) (time.Time, error) {
	db := GetDB()
	var ExpiresAt time.Time
	rows := db.QueryRow(c, "select expires_at from tokens where token=$1", token)
	err := rows.Scan(&ExpiresAt)
	if err != nil {
		return time.Time{}, err
	}
	return ExpiresAt, nil
}
