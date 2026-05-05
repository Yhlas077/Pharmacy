package models

type Pharmacy_medicines struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	New_price   float64 `json:"new_price"`
	Category_id int     `json:"category_id"`
}

type PharmacyMedicineErrorResponse struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}