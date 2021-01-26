package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/beslnet/banking/service"
)

type Customer struct {
	Name    string `json:"fullname" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zipcode"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	// customers := []Customer{
	// 	{Name: "Benjamín", City: "Santiago", Zipcode: "9849494"},
	// 	{Name: "Javier", City: "Las Condes", Zipcode: "4324324"},
	// }

	customers, _ := ch.service.GetAllCustomer()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}
