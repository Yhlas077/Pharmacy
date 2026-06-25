package repositories

import (
	"context"
	"strconv"

	"github.com/yhlas/basic-pharmacy/internal/models"
)

type PharmacyMedicineFilter struct {
	Limit  int
	Offset int
	Search string
}

func LenStrpharmacymedicine(l []any) string {
	return strconv.Itoa(len(l))
}

// GET
func PharmacyMedicineList(c context.Context, f PharmacyMedicineFilter) ([]models.PharmacyMedicines, error) {

	db := GetDB()
	if f.Limit == 0 {
		f.Limit = 10
	}
	sqlWhere := ` `
	sqlArgs := []any{f.Limit, f.Offset}
	if f.Search != "" {
		sqlArgs = append(sqlArgs, "%"+f.Search+"%")
		sqlWhere += `and name ilike $3`
	}

	rows, err := db.Query(c, `select id, name, description, price, new_price, category_id
		from pharmacy_medicines
			where 1=1 `+sqlWhere+`
		limit $1 offset $2`, sqlArgs...)
	if err != nil {
		return nil, err
	}

	list := []models.PharmacyMedicines{}

	defer rows.Close()

	for rows.Next() {
		item := models.PharmacyMedicines{}
		err := rows.Scan(&item.ID, &item.Name, &item.Description, &item.Price, &item.NewPrice, &item.CategoryId)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, nil
}

// POST /users // repository
func PharmacyMedicineCreate(c context.Context, name string, description string, price int, newPrice int, categoryID int) error {

	_, err := GetDB().Exec(context.Background(),
		"INSERT INTO pharmacy_medicines(name, description, price, new_price, category_id) VALUES ($1,$2,$3,$4,$5)",
		name, description, price, newPrice, categoryID,
	)
	if err != nil {
		return err
	}
	return nil
}

func PharmacyMedicineDelete(c context.Context, id int) error {
	db := GetDB()

	_, err := db.Exec(c,
		`DELETE FROM pharmacy_medicines WHERE id=$1`,
		id,
	)

	return err
}

func GetPharmacyMedicine(c context.Context, id int) (models.PharmacyMedicinesResponse, error) {
	db := GetDB()
	var req models.PharmacyMedicinesResponse
	rows := db.QueryRow(c, "select  id, name, description, price, new_price, category_id from pharmacy_medicines where id=$1", id)
	err := rows.Scan(&req.Id, &req.Name, &req.Description, &req.Price, &req.NewPrice, &req.CategoryId)
	if err != nil {
		return models.PharmacyMedicinesResponse{}, err
	}
	return req, nil
}

func PharmacyMedicineUpdate(c context.Context, id int, req models.PharmacyMedicinesCreateRequest) error {
	db := GetDB()

	_, err := db.Exec(c,
		`UPDATE pharmacy_medicines
		 SET name=$1, description=$2, price=$3, new_price=$4, category_id=$5 
		 WHERE id=$6`,
		req.Name, req.Description, req.Price, req.NewPrice, req.CategoryId, id,
	)

	return err
}
