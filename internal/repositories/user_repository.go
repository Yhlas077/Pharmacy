package repositories

import (
	"context"
	"strconv"

	"github.com/yhlas/basic-pharmacy/internal/models"
	"github.com/yhlas/basic-pharmacy/internal/utils"
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

func UserList(c context.Context, f UserFilter, moreArg ...int) ([]models.User, error) {
	db := utils.GetDB()
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
		rows.Scan(&item.ID, &item.Name, &item.Email, &item.Password, &item.Role)
		list = append(list, item)
	}
	return list, nil
}
