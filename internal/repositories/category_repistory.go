package repositories

import (
	"context"

	"github.com/yhlas/basic-pharmacy/internal/models"
	"github.com/yhlas/basic-pharmacy/internal/utils"
)

type CategoryFilter struct {
	Limit  int
	Offset int
}

func CategoryList(c context.Context, f CategoryFilter, moreArg ...int) ([]models.Categories, error) {
	db := utils.GetDB()
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
		rows.Scan(&item.ID, &item.Name)
		list = append(list, item)
	}
	return list, nil
}
