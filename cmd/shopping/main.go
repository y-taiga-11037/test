package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

type Products struct {
	ShoppingProductId string `json:"shopping_product_id"`
	ShoppingId        string `json:"shopping_id"`
	ProductName       string `json:"product_name"`
	Price             int    `json:"price"`
	Quantity          int    `json:"quantity"`
	PurchaseFlag      int    `json:"purchase_flag"`
}

func GetProductTableHandler(w http.ResponseWriter, r *http.Request) {
	// Read the env file
	err := godotenv.Load("/home/y-taiga/mdtd_bootcamp/env/dev.env")
	if err != nil {
		log.Fatal(err)
	}
	// Log in to mysql
	db, err := sql.Open("mysql", os.Getenv("DB_ROLE")+":"+os.Getenv("DB_PASSWORD")+"@/"+os.Getenv("DB_NAME"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// Execute a SELECT
	rows, err := db.Query("SELECT * FROM product")
	if err != nil {
		log.Fatal(err)
	}
	// Define the slice
	ProductTable := make([]Products, 0)
	for rows.Next() {
		var p Products
		// Map the SELECT result to a structure
		err = rows.Scan(&p.ShoppingProductId, &p.ShoppingId, &p.ProductName, &p.Price, &p.Quantity, &p.PurchaseFlag)
		if err != nil {
			log.Fatal(err)
		}
		// Add an element to the slice
		ProductTable = append(ProductTable, p)
	}
	// Encode to json
	res, err := json.Marshal(ProductTable)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(string(res))
	// Return json as a response
	w.Write(res)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/shopping", GetProductTableHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}
