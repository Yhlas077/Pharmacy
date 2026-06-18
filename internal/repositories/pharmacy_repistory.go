package repositories

import (
	"context"
	"strconv"

	"github.com/yhlas/basic-pharmacy/internal/models"
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

	db := GetDB()
	if f.Limit == 0 {
		f.Limit = 10
	}
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
		err := rows.Scan(&item.ID, &item.Name, &item.Address, &item.PharmacyHours)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, nil
}

// POST /Pharmacies // repository
func PharmacyCreate(c context.Context, name string, address string, hours int) error {

	_, err := GetDB().Exec(context.Background(),
		"INSERT INTO pharmacies(name, address, pharmacy_hours) VALUES ($1,$2,$3)",
		name, address, hours,
	)
	if err != nil {
		return err
	}
	return nil
}

func PharmacyDelete(c context.Context, id int) error {
	db := GetDB()

	_, err := db.Exec(c,
		`DELETE FROM pharmacies WHERE id=$1`,
		id,
	)

	return err
}

func PharmacyUpdate(c context.Context, id int, req models.PharmacyCreateRequest) error {
	db := GetDB()

	_, err := db.Exec(c,
		`UPDATE pharmacies 
		 SET name=$1, address=$2, pharmacy_hours=$3
		 WHERE id=$4`,
		req.Name, req.Address, req.PharmacyHours, id,
	)

	return err
}

func GetPharmacy(c context.Context, id int) (models.PharmacyResponse, error) {
	db := GetDB()
	var req models.PharmacyResponse
	rows := db.QueryRow(c, "select  id, name, address, pharmacy_hours from pharmacies where id=$1", id)
	err := rows.Scan(&req.Id, &req.Name, &req.Address, &req.PharmacyHours)
	if err != nil {
		return models.PharmacyResponse{}, err
	}
	return req, nil
}
