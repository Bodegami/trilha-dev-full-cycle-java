package main

import (
	"database/sql"
	db2 "github.com/bodegami/trilha-dev-full-cycle-java/arquitetura-hexagonal/adapter/db"
	"github.com/bodegami/trilha-dev-full-cycle-java/arquitetura-hexagonal/application"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var db *sql.DB
	db, _ = sql.Open("sqlite3", "db.sqlite")
	productDbAdapter := db2.NewProductDb(db)
	productService := application.NewProductService(productDbAdapter)

	product, _ := productService.Create("Product Example", 30)

	productService.Enable(product)
}
