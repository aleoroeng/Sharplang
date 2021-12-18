package model

import "fmt"

type Product struct {
	ID           int
	Description  string
	Name         string
	Price        int
	Stock        int
	TimesUpdated int
}

func (p Product) String() string {
	return fmt.Sprintf("id: %d, description: %v, name: %v, price: %d, stock: %d, timesUpdated: %d", p.ID, p.Description, p.Name, p.Price,
		p.Stock, p.TimesUpdated)
}
