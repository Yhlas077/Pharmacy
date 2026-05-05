package repositories

import (
	"context"

	"github.com/yhlas/basic-pharmacy/internal/models"
	"github.com/yhlas/basic-pharmacy/internal/utils"
)

type PharmacyMedicineFilter struct {
	Limit  int
	Offset int
}

func Pharmacy_Medicine_List(c context.Context, f PharmacyMedicineFilter, moreArg ...int) ([]models.Pharmacy_medicines, error) {
	db := utils.GetDB()
	sqlWhere := ` `
	sqlArgs := []any{f.Limit, f.Offset}

	rows, err := db.Query(c, `select id, name, description, price, new_price, category_id
		from pharmacy_medicines
			where 1=1 `+sqlWhere+`
		limit $1 offset $2`, sqlArgs...)
	if err != nil {
		return nil, err
	}

	list := []models.Pharmacy_medicines{}

	defer rows.Close()

	for rows.Next() {
		item := models.Pharmacy_medicines{}
		rows.Scan(&item.ID, &item.Name, &item.Description, &item.Price, &item.New_price, &item.Category_id)
		list = append(list, item)
	}
	return list, nil
}
