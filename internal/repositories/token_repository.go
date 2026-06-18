package repositories

import (
	"context"
	"fmt"
)

func InsertToken(c context.Context, id int, token string) error {
	_, err := GetDB().Exec(c, "INSERT into tokens(user_id, token) values ($1, $2)", id, token)
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

func InsertUser(c context.Context, name string, email string, password string) error {
	_, err := GetDB().Exec(c,
		"INSERT INTO users(name, email, password, role) VALUES ($1,$2,$3,$4)",
		name, email, password, "user",
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
