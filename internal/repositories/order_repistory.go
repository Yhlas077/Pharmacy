package repositories

import (
	"context"

	"github.com/yhlas/basic-pharmacy/internal/models"
	"github.com/yhlas/basic-pharmacy/internal/utils"
)

type OrderFilter struct {
	Limit  int
	Offset int
}

func OrderList(c context.Context, f OrderFilter, moreArg ...int) ([]models.Orders, error) {
	db := utils.GetDB()
	sqlWhere := ` `
	sqlArgs := []any{f.Limit, f.Offset}
	rows, err := db.Query(c, `select id, name, price, description
		from orders
			where 1=1 `+sqlWhere+`
		limit $1 offset $2`, sqlArgs...)
	if err != nil {
		return nil, err
	}

	list := []models.Orders{}

	defer rows.Close()

	for rows.Next() {
		item := models.Orders{}
		rows.Scan(&item.ID, &item.Name, &item.Price, &item.Description)
		list = append(list, item)
	}
	return list, nil
}
