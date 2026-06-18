package services

import (
	"context"

	"github.com/yhlas/basic-pharmacy/internal/models"
	"github.com/yhlas/basic-pharmacy/internal/repositories"
)

func PharmacyMedicineListService(c context.Context, filter repositories.PharmacyMedicineFilter) (any, error) {
	return repositories.PharmacyMedicineList(c, filter)
}
func CreatePharmacyMedicineService(c context.Context, name string, description string, price int, newprice int, categoryid int, token string) error {
	err := repositories.PharmacyMedicineCreate(c, name, description, price, newprice, categoryid)
	if err != nil {
		return err
	}
	return nil
}
func DeletePharmacyMedicineService(c context.Context, id int) error {
	return repositories.PharmacyMedicineDelete(c, id)
}
func UpdatePharmacyMedicineService(c context.Context, id int, req models.PharmacyMedicinesCreateRequest) error {
	return repositories.PharmacyMedicineUpdate(c, id, req)
}
func GetPharmacyMedicineServices(c context.Context, id int) (models.OrderResponse, error) {
	return repositories.GetOrder(c, id)
}
