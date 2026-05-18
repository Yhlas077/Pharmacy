package models

type Pharmacies struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Address        string `json:"address"`
	Pharmacy_hours int    `json:"pharmacy_hours"`
}
