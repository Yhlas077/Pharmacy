package services

import (
	"context"

	"github.com/yhlas/basic-pharmacy/internal/models"
	"github.com/yhlas/basic-pharmacy/internal/repositories"
)

func PharmacyListService(c context.Context, filter repositories.PharmacyFilter) (any, error) {
	return repositories.PharmacyList(c, filter)
}
func CreatePharmacyService(c context.Context, name string, address string, hours int) error {
	return repositories.PharmacyCreate(c, name, address, hours)
}
func DeletePharmacyService(c context.Context, categoryid int) error {
	return repositories.PharmacyDelete(c, categoryid)
}
func UpdatePharmacyService(c context.Context, id int, req models.PharmacyCreateRequest) error {
	return repositories.PharmacyUpdate(c, id, req)
}
func GetPharmacyService(c context.Context, id int) (models.PharmacyResponse, error) {
	return repositories.GetPharmacy(c, id)
}