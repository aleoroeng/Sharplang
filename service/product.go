package service

import (
	"database/sql"
	"fmt"
	"log"
	"sharplang/src/model"
)

type Product struct {
	model.Product
}

func NewProduct() *Product {
	var product Product
	return &product
}

func (p *Product) GetAllProducts(tx *sql.Tx) (products []*model.Product) {
	defer tx.Rollback()
	rows, rowsErr := tx.Query("SELECT * FROM product")
	defer rows.Close()
	if rowsErr != nil {
		log.Fatal(rowsErr)
	}

	for rows.Next() {
		var product model.Product
		err := rows.Scan(&product.ID, &product.Description, &product.Name, &product.Price, &product.Stock, &product.TimesUpdated)
		if err != nil {
			log.Fatal(err)
		}
		products = append(products, &product)
		fmt.Printf("id: %d desc: %s name: %s price: %d stock: %d timesUpdated: %d\n", product.ID, product.Description, product.Name, product.Price, product.Stock, product.TimesUpdated)
	}
	err := tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
	return
}

func (p *Product) GetProductById(tx *sql.Tx, id int) *model.Product {
	defer tx.Rollback()
	rows, rowsErr := tx.Query("SELECT * FROM product WHERE id = $1", id)

	if rowsErr != nil {
		fmt.Println(rowsErr)
		return nil
	}

	var product model.Product
	for rows.Next() {
		rows.Scan(&product.ID, &product.Description, &product.Name, &product.Price, &product.Stock, &product.TimesUpdated)
	}

	if errRows := rows.Err(); errRows != nil {
		return nil
	}
	fmt.Printf("id: %d, desc: %v, name: %v, price: %d, stock: %d, timesUpdated: %d", product.ID, product.Description, product.Name, product.Price, product.Stock, product.TimesUpdated)
	tx.Commit()
	return &product
}
