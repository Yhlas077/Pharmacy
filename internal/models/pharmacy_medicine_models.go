package models

type PharmacyMedicines struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	New_price   float64 `json:"new_price"`
	Category_id int     `json:"category_id"`
}
