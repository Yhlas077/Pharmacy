package services

import (
	"context"

	"github.com/yhlas/basic-pharmacy/internal/models"
	"github.com/yhlas/basic-pharmacy/internal/repositories"
)

func CategoryListService(c context.Context, filter repositories.CategoryFilter) (any, error) {
	return repositories.CategoryList(c, filter)
}
func CreateCategoryService(c context.Context, name string) error {
	return repositories.CategoryCreate(c, name)
}
func DeleteCategoryService(c context.Context, categoryid int) error {
	return repositories.CategoryDelete(c, categoryid)
}
func UpdateCategoryService(c context.Context, categoryid int, req models.CategoryCreateRequest) error {
	return repositories.CategoryUpdate(c, categoryid, req)
}
func GetCategoryService(c context.Context, categoryid int) (models.CategoryResponse, error) {
	return repositories.GetCategory(c, categoryid)
}
