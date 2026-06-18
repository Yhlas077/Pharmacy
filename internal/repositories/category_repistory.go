package repositories

import (
	"context"
	"strconv"

	"github.com/yhlas/basic-pharmacy/internal/models"
)

type CategoryFilter struct {
	Limit  int
	Offset int
	Search string
}

func LenStrcategory(l []any) string {
	return strconv.Itoa(len(l))
}

// GET
func CategoryList(c context.Context, f CategoryFilter) ([]models.Category, error) {

	db := GetDB()
	if f.Limit == 0 {
		f.Limit = 10
	}
	sqlWhere := ` `
	sqlArgs := []any{f.Limit, f.Offset}
	if f.Search != "" {
		sqlArgs = append(sqlArgs, f.Search)
		sqlWhere += `and (name ilike '%$` + LenStrcategory(sqlArgs) + `%')`
	}

	rows, err := db.Query(c, `select id, name
		from categories
			where 1=1 `+sqlWhere+`
		limit $1 offset $2`, sqlArgs...)
	if err != nil {
		return nil, err
	}

	list := []models.Category{}

	defer rows.Close()

	for rows.Next() {
		item := models.Category{}
		err := rows.Scan(&item.ID, &item.Name)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, nil
}

func GetCategory(c context.Context, categoryid int) (models.CategoryResponse, error) {
	db := GetDB()
	var req models.CategoryResponse
	rows := db.QueryRow(c, "select  id, name from categories where id=$1", categoryid)
	err := rows.Scan(&req.ID, &req.Name)
	if err != nil {
		return models.CategoryResponse{}, err
	}
	return req, nil
}

// POST /Category // repository
func CategoryCreate(c context.Context, name string) error {

	_, err := GetDB().Exec(context.Background(),
		"INSERT INTO categories(name) VALUES ($1)", name,
	)
	if err != nil {
		return err
	}
	return nil
}

func CategoryDelete(c context.Context, id int) error {
	db := GetDB()

	_, err := db.Exec(c,
		`DELETE FROM categories WHERE id=$1`,
		id,
	)

	return err

}

func CategoryUpdate(c context.Context, id int, req models.CategoryCreateRequest) error {
	db := GetDB()

	_, err := db.Exec(c,
		`UPDATE categories 
		 SET name=$1
		 WHERE id=$2`,
		req.Name, id,
	)

	return err
}
