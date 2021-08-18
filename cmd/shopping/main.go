package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

type ProductStructer struct {
	ProductName string `json:"product_name"`
	Quantity    int    `json:"quantity"`
	Price       int    `json:"price"`
}

type ResponseStructer struct {
	ShoppingId  int               `json:"shopping_id"`
	ShoppingDay string            `json:"shopping_day"`
	Products    []ProductStructer `json:"products"`
}

func GetJoinTableHandler(w http.ResponseWriter, r *http.Request) {

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
	rows, err := db.Query("SELECT P.shopping_id, S.shopping_day, P.product_name, P.quantity, P.price FROM shopping AS S LEFT OUTER JOIN product AS P ON P.shopping_id = S.shopping_id;")
	if err != nil {
		log.Fatal(err)
	}

	ResponseSlice := make([]ResponseStructer, 0)
	var counter int = 0
	var days string

	var container ResponseStructer
	var subcontainer ResponseStructer

	for rows.Next() {
		var PS ProductStructer

		// Map the SELECT result to a structure
		err = rows.Scan(&subcontainer.ShoppingId, &subcontainer.ShoppingDay, &PS.ProductName, &PS.Quantity, &PS.Price)

		if err != nil {
			log.Fatal(err)
		}

		// Add an element to the slice

		if counter == subcontainer.ShoppingId {

			container.Products = append(container.Products, PS)

			fmt.Println("counter = ", counter)
			fmt.Println("days = ", days)
			fmt.Println("container.Products = ", container.Products)

		} else if counter == 0 {

			fmt.Println("counter = ", counter)

			container.Products = append(container.Products, PS)

			counter = subcontainer.ShoppingId
			days = subcontainer.ShoppingDay

			fmt.Println("container.Products = ", container.Products)
			fmt.Println("counter = ", counter)
			fmt.Println("days = ", days)

		} else {

			fmt.Println("counter = ", counter)
			fmt.Println("days = ", days)

			container = ResponseStructer{ShoppingId: counter, ShoppingDay: days, Products: container.Products}
			ResponseSlice = append(ResponseSlice, container)

			container = ResponseStructer{}

			container.Products = append(container.Products, PS)

			counter = subcontainer.ShoppingId
			days = subcontainer.ShoppingDay

			fmt.Println("ResponseSlice = ", ResponseSlice)
			fmt.Println("counter = ", counter)
			fmt.Println("days = ", days)
			subcontainer = ResponseStructer{}

		}

	}

	container = ResponseStructer{ShoppingId: counter, ShoppingDay: days, Products: container.Products}
	ResponseSlice = append(ResponseSlice, container)

	fmt.Println("ResponseSlice = ", ResponseSlice)
	fmt.Println("counter = ", counter)
	fmt.Println("days = ", days)

	// Encode to json
	res, err := json.Marshal(ResponseSlice)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(string(res))

	// Return json as a response
	w.Write(res)
}

func PostCreateHandler(w http.ResponseWriter, r *http.Request) {

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

	InsertStatement := "INSERT INTO shopping (shopping_day) VALUES (?);"
	stmt, err := db.Prepare(InsertStatement)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(r.Body)
	fmt.Println(string(body))
	if err != nil && err != io.EOF {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	var FirstInsert ResponseStructer
	err = json.Unmarshal(body, &FirstInsert)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal(err)
		return
	}

	fmt.Printf("%v\n", FirstInsert)

	result, err := stmt.Exec(&FirstInsert.ShoppingDay)

	stmt.Close()

	if err != nil {
		log.Fatal(err)
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(lastInsertID)

	var SecondInsert ResponseStructer
	err = json.Unmarshal(body, &SecondInsert)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal(err)
		return
	}

	fmt.Printf("%v\n", SecondInsert)

	InsertStatement2 := "INSERT INTO product (shopping_id, product_name, price, quantity) VALUES (?, ?, ?, ?);"
	stmt2, err := db.Prepare(InsertStatement2)
	if err != nil {
		log.Fatal(err)
	}

	defer stmt2.Close()

	for i := range SecondInsert.Products {

		result2, err := stmt2.Exec(lastInsertID, &SecondInsert.Products[i].ProductName, &SecondInsert.Products[i].Price, &SecondInsert.Products[i].Quantity)
		if err != nil {
			log.Fatal(err)
		}
		lastInsertID2, err := result2.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(lastInsertID2)

	}

}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/api/shopping2", GetJoinTableHandler)
	r.HandleFunc("/api/shopping3", PostCreateHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}
