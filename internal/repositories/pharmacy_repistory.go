package repositories

import (
	"context"

	"github.com/yhlas/basic-pharmacy/internal/models"
	"github.com/yhlas/basic-pharmacy/internal/utils"
)

type PharmaciesFilter struct {
	Limit  int
	Offset int
}

func PharmaciesList(c context.Context, f PharmaciesFilter, moreArg ...int) ([]models.Pharmacies, error) {
	db := utils.GetDB()
	sqlWhere := ` `
	sqlArgs := []any{f.Limit, f.Offset}

	rows, err := db.Query(c, `select id, name, address, pharmacy_hours
		from pharmacies
			where 1=1 `+sqlWhere+`
		limit $1 offset $2`, sqlArgs...)
	if err != nil {
		return nil, err
	}

	list := []models.Pharmacies{}

	defer rows.Close()

	for rows.Next() {
		item := models.Pharmacies{}
		rows.Scan(&item.ID, &item.Name, &item.Address, &item.Pharmacy_hours)
		list = append(list, item)
	}
	return list, nil
}
