package repositories

import (
	"context"
	"fmt"
)

func ShowToken(c context.Context, token string) {

}

func InsertToken(c context.Context, id int, token string) {
	GetDB().Exec(c, "INSERT into tokens(user_id, token) values ($1, $2)", id, token)

}

func GetToken(c context.Context, token string, id int) {
	row := GetDB().QueryRow(c, `select token from tokens where id =$1`, id)
	row.Scan(&token)
}

func DeleteToken(c context.Context, token string) {
	var id int
	row := GetDB().QueryRow(c, `select user_id from tokens where token = $1`, token)
	row.Scan(&id)
	GetDB().Exec(c, `DELETE FROM tokens WHERE user_id=$1`, id)

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
