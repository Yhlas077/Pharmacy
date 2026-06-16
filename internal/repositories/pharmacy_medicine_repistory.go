package repositories

import (
	"context"

	"github.com/yhlas/basic-pharmacy/internal/models"
)

type PharmacyMedicineFilter struct {
	Limit  int
	Offset int
	Search string
}

// GET
func PharmacyMedicineList(c context.Context, f PharmacyMedicineFilter) ([]models.PharmacyMedicines, error) {

	db := GetDB()
	sqlWhere := ` `
	sqlArgs := []any{f.Limit, f.Offset}

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
func PharmacyMedicineCreate(c context.Context, pharmacy_medicine models.PharmacyMedicines) (models.PharmacyMedicines, error) {

	_, err := GetDB().Exec(context.Background(),
		"INSERT INTO pharmacy_medicines(id, name, description, price, new_price, category_id) VALUES ($1,$2,$3,$4,$5,$6)",
		pharmacy_medicine.ID, pharmacy_medicine.Name, pharmacy_medicine.Description, pharmacy_medicine.Price, pharmacy_medicine.NewPrice, pharmacy_medicine.CategoryId,
	)
	if err != nil {
		return models.PharmacyMedicines{}, err
	}
	return pharmacy_medicine, nil
}

func PharmacyMedicineDelete(c context.Context, id int) error {
	db := GetDB()

	_, err := db.Exec(c,
		`DELETE FROM pharmacy_medicines WHERE id=$1`,
		id,
	)

	return err
}

func PharmacyMedicineUpdate(c context.Context, id int, req models.PharmacyMedicines) error {
	db := GetDB()

	_, err := db.Exec(c,
		`UPDATE pharmacy_medicines
		 SET name=$1, description=$2, price=$3, new_price=$4, category_id=$5 
		 WHERE id=$6`,
		req.Name, req.Description, req.Price, req.NewPrice, req.CategoryId, id,
	)

	return err
}
