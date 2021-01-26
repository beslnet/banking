// Package domain provides ...
package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "Benjamín", "Santiago", "43243", "2000-01-01", "1"},
		{"1002", "Javier", "Las Condes", "32321", "2000-02-01", "1"},
	}

	return CustomerRepositoryStub{customers: customers}
}
