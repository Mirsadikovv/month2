package main

import (
	"database/sql"
	"fmt"
	"log"
	"market-management/config"
	"market-management/model"
	"market-management/postgres"
)

func main() {
	cfg := config.Load()
	postgresDB, err := postgres.NewPostgres(cfg)
	if err != nil {
		log.Panic("failed to connect DB", err)
	}
	defer postgresDB.Close()

	product := model.Products{
		Product_name:     "phone",
		Product_price:    1100,
		Product_quantity: 3,
	}
	product_categorie := model.Categories{
		Category_name: "technology",
		Product_id:    1,
	}

	product2 := model.Products{
		Product_name:     "phone2",
		Product_price:    11001,
		Product_quantity: 4,
	}
	product_categorie2 := model.Categories{
		Category_name: "technology22",
		Product_id:    2,
	}
	idd := createProduct(postgresDB, product, product_categorie)
	idd2 := createProduct(postgresDB, product2, product_categorie2)
	fmt.Println(selectProduct(postgresDB, idd), idd)
	fmt.Println(selectProduct(postgresDB, idd2), idd2)

}

func createProduct(db *sql.DB, product model.Products, category model.Categories) int {
	query := `add_product($1,$2,$3,$4)`
	var id int
	db.QueryRow(query, product.Product_name, product.Product_price, product.Product_quantity, category.Category_name).Scan(&id)
	return id
}

func selectProduct(db *sql.DB, id int) model.Prod_and_cat {
	var product model.Prod_and_cat
	query := `get_product_info($1)`
	row := db.QueryRow(query, id)
	row.Scan(&product.Id, &product.Product_name, &product.Product_price, &product.Product_quantity, &product.Product_categorie)
	return product
}
