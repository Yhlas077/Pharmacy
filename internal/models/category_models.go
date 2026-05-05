package models

type Categories struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CategoriesErrorResponse struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}