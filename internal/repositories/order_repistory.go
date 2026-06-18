package repositories

import (
	"context"
	"strconv"

	"github.com/yhlas/basic-pharmacy/internal/models"
)

type OrdersFilter struct {
	Limit      int
	Offset     int
	UserId     int
	PharmacyID int
	Search     string
}

func LenStrorder(l []any) string {
	return strconv.Itoa(len(l))
}

// GET
func OrdersList(c context.Context, f OrdersFilter) ([]models.Orders, error) {

	db := GetDB()
	if f.Limit == 0 {
		f.Limit = 10
	}
	sqlWhere := ` `
	sqlArgs := []any{f.Limit, f.Offset}
	if f.Search != "" {
		sqlArgs = append(sqlArgs, f.Search)
		sqlWhere += `and (name ilike '%$` + LenStrorder(sqlArgs) + `%')`
	}

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
func OrdersCreate(c context.Context, name string, price float64, description string) error {

	_, err := GetDB().Exec(context.Background(),
		"INSERT INTO orders(name, price, description) VALUES ($1,$2,$3)",
		name, price, description,
	)
	if err != nil {
		return err
	}
	return nil
}

func OrdersDelete(c context.Context, id int) error {
	db := GetDB()

	_, err := db.Exec(c,
		`DELETE FROM orders WHERE id=$1`,
		id,
	)

	return err
}

func OrdersUpdate(c context.Context, id int, req models.OrderCreateRequest) error {
	db := GetDB()

	_, err := db.Exec(c,
		`UPDATE orders 
		 SET name=$1, price=$2, description=$3
		 WHERE id=$4`,
		req.Name, req.Price, req.Description, id,
	)

	return err
}

func GetOrder(c context.Context, id int) (models.OrderResponse, error) {
	db := GetDB()
	var req models.OrderResponse
	rows := db.QueryRow(c, "select  id, name, price, description from orders where id=$1", id)
	err := rows.Scan(&req.ID, &req.Name, &req.Price, &req.Description)
	if err != nil {
		return models.OrderResponse{}, err
	}
	return req, nil
}
