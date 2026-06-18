package models

type PharmacyMedicines struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	NewPrice    float64 `json:"new_price"`
	CategoryId  int     `json:"category_id"`
}

type PharmacyMedicinesResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	NewPrice    int    `json:"newprice"`
	CategoryId  int    `json:"categoryid"`
}

type PharmacyMedicinesCreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	NewPrice    int    `json:"newprice"`
	CategoryId  int    `json:"categoryid"`
}