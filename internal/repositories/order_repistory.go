package repositories

import (
	"context"

	"github.com/yhlas/basic-pharmacy/internal/models"
)

type OrdersFilter struct {
	Limit      int
	Offset     int
	UserId     int
	PharmacyID int
}

// GET
func OrdersList(c context.Context, f OrdersFilter) ([]models.Orders, error) {

	db := GetDB()
	sqlWhere := ` `
	sqlArgs := []any{f.Limit, f.Offset}

	// TODO: implement FILTERS BY LOGIC (DB)
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
		err := rows.Scan(&item.ID, &item.Name, &item.Price, &item.Description)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, nil
}

// POST /orders // repository
func OrdersCreate(c context.Context, Orders models.Orders) (models.Orders, error) {

	_, err := GetDB().Exec(context.Background(),
		"INSERT INTO orders(id, name, price, description) VALUES ($1,$2,$3,$4)",
		Orders.ID, Orders.Name, Orders.Price, Orders.Description,
	)
	if err != nil {
		return models.Orders{}, err
	}
	return Orders, nil
}

func OrdersDelete(c context.Context, id int) error {
	db := GetDB()

	_, err := db.Exec(c,
		`DELETE FROM orders WHERE id=$1`,
		id,
	)

	return err
}

func OrdersUpdate(c context.Context, id int, req models.Orders) error {
	db := GetDB()

	_, err := db.Exec(c,
		`UPDATE orders 
		 SET name=$1, price=$2, description=$3
		 WHERE id=$4`,
		req.Name, req.Price, req.Description, id,
	)

	return err
}
