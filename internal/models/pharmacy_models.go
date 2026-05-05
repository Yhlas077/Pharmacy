package models

type Pharmacies struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Address        string `json:"address"`
	Pharmacy_hours int    `json:"pharmacy_hours"`
}

type PharmacyErrorResponse struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}