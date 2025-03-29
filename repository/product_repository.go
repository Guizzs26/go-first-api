package repository

import (
	"database/sql"
	"fmt"

	"github.com/Guizzs26/go-first-api/models"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) *ProductRepository {
	return &ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]models.Product, error) {
	query := "SELECT * FROM products LIMIT 10"

	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []models.Product{}, err
	}

	var productList []models.Product
	var productObj models.Product

	for rows.Next() {
		err = rows.Scan(
			&productObj.ID,
			&productObj.Name,
			&productObj.Price,
		)
		if err != nil {
			fmt.Println(err)
			return []models.Product{}, err
		}

		productList = append(productList, productObj)
	}
	rows.Close()

	return productList, nil
}
