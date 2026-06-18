package models

type Pharmacies struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	Address       string  `json:"address"`
	PharmacyHours int     `json:"pharmacy_hours"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
}

type PharmacyResponse struct {
	Id            int     `json:"id"`
	Name          string  `json:"name"`
	Address       string  `json:"address"`
	PharmacyHours int     `json:"pharmacy_hours"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
}

type PharmacyCreateRequest struct {
	Name          string  `json:"name"`
	Address       string  `json:"address"`
	PharmacyHours int     `json:"pharmacy_hours"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
}
type NearPharmacies struct {
	Name string `json:"name"`
}
