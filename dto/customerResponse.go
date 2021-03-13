package dto

type CustomerResponse struct {
	Id          string `json:"customer_id"`
	Name        string `json:"fullname"`
	City        string `json:"city"`
	Zipcode     string `json:"zip_code"`
	DateOfBirth string `db:"date_of_birth"`
	Status      string `json:"status"`
}
