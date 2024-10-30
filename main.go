package main

import (
	"database/sql"

	db2 "github.com/higorrsc/fc-hrsc-hexagonal-architecture/adapters/db"
	"github.com/higorrsc/fc-hrsc-hexagonal-architecture/application"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "db.sqlite3")
	defer db.Close()

	productDbAdapter := db2.NewProductDb(db)
	productService := application.NewProductService(productDbAdapter)

	product, _ := productService.Create("Product Example", 30)
	productService.Enable(product)

}
