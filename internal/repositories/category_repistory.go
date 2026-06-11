package repositories

import (
	"context"

	"github.com/yhlas/basic-pharmacy/internal/models"
)

type CategoryFilter struct {
	Limit  int
	Offset int
}

// GET
func CategoryList(c context.Context, f CategoryFilter) ([]models.Categories, error) {

	db := GetDB()
	sqlWhere := ` `
	sqlArgs := []any{f.Limit, f.Offset}

	rows, err := db.Query(c, `select id, name
		from categories
			where 1=1 `+sqlWhere+`
		limit $1 offset $2`, sqlArgs...)
	if err != nil {
		return nil, err
	}

	list := []models.Categories{}

	defer rows.Close()

	for rows.Next() {
		item := models.Categories{}
		err := rows.Scan(&item.ID, &item.Name)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, nil
}

// POST /Category // repository
func CategoryCreate(c context.Context, Category models.Categories) (models.Categories, error) {

	_, err := GetDB().Exec(context.Background(),
		"INSERT INTO categories(id, name) VALUES ($1,$2)",
		Category.ID, Category.Name,
	)
	if err != nil {
		return models.Categories{}, err
	}
	return Category, nil
}

func CategoryDelete(c context.Context, id int) error {
	db := GetDB()

	_, err := db.Exec(c,
		`DELETE FROM categories WHERE id=$1`,
		id,
	)

	return err

}

func CategoryUpdate(c context.Context, id int, req models.Categories) error {
	db := GetDB()

	_, err := db.Exec(c,
		`UPDATE categories 
		 SET name=$1
		 WHERE id=$2`,
		req.Name, id,
	)

	return err
}
