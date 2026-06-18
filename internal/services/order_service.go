package services

import (
	"context"

	"github.com/yhlas/basic-pharmacy/internal/models"
	"github.com/yhlas/basic-pharmacy/internal/repositories"
)

func OrderListService(c context.Context, filter repositories.OrdersFilter) (any, error) {
	return repositories.OrdersList(c, filter)
}
func CreateOrderService(c context.Context, name string, price float64, description string) error {
	return repositories.OrdersCreate(c, name, price, description)
}
func DeleteOrderService(c context.Context, id int) error {
	return repositories.OrdersDelete(c, id)
}
func UpdateOrderService(c context.Context, id int, req models.OrderCreateRequest) error {
	return repositories.OrdersUpdate(c, id, req)
}
func GetOrderServices(c context.Context, id int) (models.OrderResponse, error) {
	return repositories.GetOrder(c, id)
}
