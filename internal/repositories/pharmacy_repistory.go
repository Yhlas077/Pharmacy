package repositories

import (
	"context"
	"strconv"

	"github.com/yhlas/basic-pharmacy/internal/models"
	"github.com/yhlas/basic-pharmacy/internal/utils"
)

type PharmacyFilter struct {
	Limit  int
	Offset int
	Search string
}

func LengthStr(l []any) string {
	return strconv.Itoa(len(l))
}

// GET
func PharmacyList(c context.Context, f PharmacyFilter) ([]models.Pharmacies, error) {

	db := utils.GetDB()
	sqlWhere := ` `
	sqlArgs := []any{f.Limit, f.Offset}

	if f.Search != "" {
		sqlArgs = append(sqlArgs, f.Search)
		sqlWhere += ` and (first_name ilike '%$` + LengthStr(sqlArgs) + `%')`

	}

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
		err := rows.Scan(&item.ID, &item.Name, &item.Address, &item.Pharmacy_hours)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, nil
}

// POST /Pharmacies // repository
func PharmacyCreate(c context.Context, Pharmacy models.Pharmacies) (models.Pharmacies, error) {

	_, err := utils.GetDB().Exec(context.Background(),
		"INSERT INTO pharmacies(id, name, address, pharmacy_hours) VALUES ($1,$2,$3,$4)",
		Pharmacy.ID, Pharmacy.Name, Pharmacy.Address, Pharmacy.Pharmacy_hours,
	)
	if err != nil {
		return models.Pharmacies{}, err
	}
	return Pharmacy, nil
}

func PharmacyDelete(c context.Context, id int) error {
	db := utils.GetDB()

	_, err := db.Exec(c,
		`DELETE FROM pharmacies WHERE id=$1`,
		id,
	)

	return err
}

func PharmacyUpdate(c context.Context, id int, req models.Pharmacies) error {
	db := utils.GetDB()

	_, err := db.Exec(c,
		`UPDATE pharmacies 
		 SET name=$1, address=$2, pharmacy_hours=$3
		 WHERE id=$4`,
		req.Name, req.Address, req.Pharmacy_hours, id,
	)

	return err
}
