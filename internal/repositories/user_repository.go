package repositories

import (
	"context"
	"strconv"

	"github.com/yhlas/basic-pharmacy/internal/models"
)

type UserFilter struct {
	Limit  int
	Offset int
	Search string
	Role   string
}

func LenStr(l []any) string {
	return strconv.Itoa(len(l))
}

// GET
func UserList(c context.Context, f UserFilter) ([]models.User, error) {

	db := GetDB()
	if f.Limit == 0 {
		f.Limit = 10
	}
	sqlWhere := ` `
	sqlArgs := []any{f.Limit, f.Offset}
	if f.Search != "" {
		sqlArgs = append(sqlArgs, f.Search)
		sqlWhere += ` and (first_name ilike '%$` + LenStr(sqlArgs) + `%')`

	}
	if f.Role != "" {
		sqlArgs = append(sqlArgs, f.Role)
		sqlWhere += ` and (role=$` + LenStr(sqlArgs) + `)`
	}

	rows, err := db.Query(c, `select id, name, email, password, role
		from users
			where 1=1 `+sqlWhere+`
		limit $1 offset $2`, sqlArgs...)
	if err != nil {
		return nil, err
	}

	list := []models.User{}

	defer rows.Close()

	for rows.Next() {
		item := models.User{}
		err := rows.Scan(&item.ID, &item.Name, &item.Email, &item.Password, &item.Role)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, nil
}

func UserGetByEmail(c context.Context, email string) (models.User, error) {

	db := GetDB()

	row := db.QueryRow(c, `select id, name, email, password, role
		from users
			where email=$1`, email)

	item := models.User{}

	err := row.Scan(&item.ID, &item.Name, &item.Email, &item.Password, &item.Role)
	if err != nil {
		return models.User{}, err
	}
	return item, nil
}

func UserGetByID(c context.Context, id int) (models.User, error) {

	db := GetDB()

	row := db.QueryRow(c, `select id, name, email, password, role
		from users
			where id=$1`, id)

	item := models.User{}

	err := row.Scan(&item.ID, &item.Name, &item.Email, &item.Password, &item.Role)
	if err != nil {
		return models.User{}, err
	}
	return item, nil
}

func GetUser(c context.Context, token string, hasPass bool) (models.UserResponse, error) {
	db := GetDB()
	var req models.UserResponse
	rows := db.QueryRow(c, "select  u.id, u.name, u.role, u.password, u.email from users u join tokens t on t.user_id=u.id where t.token=$1", token)
	err := rows.Scan(&req.ID, &req.Name, &req.Role, &req.Password, &req.Email)
	if !hasPass {
		req.Password = ""
	}
	if err != nil {
		return models.UserResponse{}, err
	}
	return req, nil
}

// POST /users // repository
func UserCreate(c context.Context, user models.User) (models.User, error) {

	_, err := GetDB().Exec(context.Background(),
		"INSERT INTO users(id, name, email, password, role) VALUES ($1,$2,$3,$4,$5)",
		user.ID, user.Name, user.Email, user.Password, user.Role,
	)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func UserDelete(c context.Context, id int) error {
	db := GetDB()

	_, err := db.Exec(c,
		`DELETE FROM users WHERE id=$1`,
		id,
	)

	return err
}

func UpdatePassword(c context.Context, token string, passchange models.ChangePasswordRequest) error {
	db := GetDB()

	_, err := db.Exec(c,
		`update users u set password=$1 from tokens t where t.user_id=u.id and t.token=$2`,
		passchange.NewPassword, token)

	return err
}

func UpdateUser(c context.Context, token string, req models.UserUpdateRequest) error {
	db := GetDB()

	_, err := db.Exec(c, "update users u set name=$1 from tokens t where t.user_id=u.id and t.token=$2", req.Name, token)
	if err != nil {
		return err
	}
	return nil
}