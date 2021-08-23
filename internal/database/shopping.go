package database

import (
	"encoding/json"

	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	ProductName string `json:"product_name"`
	Quantity    int    `json:"quantity"`
	Price       int    `json:"price"`
}

type Response struct {
	ShoppingId  int       `json:"shopping_id"`
	ShoppingDay string    `json:"shopping_day"`
	Products    []Product `json:"products"`
}

func GetSelectShoppingListsQuery() ([]Response, error) {

	responseSlice := make([]Response, 0)

	// Execute a SELECT
	rows, err := DB.Query("SELECT P.shopping_id, S.shopping_day, P.product_name, P.quantity, P.price FROM shopping AS S LEFT OUTER JOIN product AS P ON P.shopping_id = S.shopping_id;")
	if err != nil {
		return responseSlice, err
	}

	var id int = 0
	var day string

	var container Response
	var subcontainer Response

	for rows.Next() {
		var ps Product

		// Map the SELECT result to a structure
		err = rows.Scan(&subcontainer.ShoppingId, &subcontainer.ShoppingDay, &ps.ProductName, &ps.Quantity, &ps.Price)
		if err != nil {
			return responseSlice, err
		}

		// Add an element to the slice

		if id == subcontainer.ShoppingId {

			container.Products = append(container.Products, ps)

		} else if id == 0 {

			container.Products = append(container.Products, ps)

			id = subcontainer.ShoppingId
			day = subcontainer.ShoppingDay

		} else {

			container = Response{ShoppingId: id, ShoppingDay: day, Products: container.Products}
			responseSlice = append(responseSlice, container)

			container = Response{}

			container.Products = append(container.Products, ps)

			id = subcontainer.ShoppingId
			day = subcontainer.ShoppingDay

			subcontainer = Response{}

		}

	}
	container = Response{ShoppingId: id, ShoppingDay: day, Products: container.Products}
	responseSlice = append(responseSlice, container)

	return responseSlice, err
}

func InsertShopping(body []byte) (int64, error) {

	var firstInsert Response
	err := json.Unmarshal(body, &firstInsert)
	if err != nil {
		return 0, err
	}

	insertStatement := "INSERT INTO shopping (shopping_day) VALUES (?);"
	stmt, err := DB.Prepare(insertStatement)
	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(&firstInsert.ShoppingDay)
	if err != nil {
		return 0, err
	}

	stmt.Close()

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastInsertID, err
}

func GetPostQuery(lastInsertID int64) ([]Response, error) {

	responseSlice := make([]Response, 0)

	// Execute a SELECT
	rows, err := DB.Query("SELECT P.shopping_id, S.shopping_day, P.product_name, P.quantity, P.price FROM shopping AS S LEFT OUTER JOIN product AS P ON P.shopping_id = S.shopping_id WHERE P.shopping_id = CONCAT(?, '%')", lastInsertID)
	if err != nil {
		return responseSlice, err
	}

	var counter int = 0
	var day string
	var container Response
	var subcontainer Response

	for rows.Next() {
		var ps Product

		// Map the SELECT result to a structure
		err = rows.Scan(&subcontainer.ShoppingId, &subcontainer.ShoppingDay, &ps.ProductName, &ps.Quantity, &ps.Price)
		if err != nil {
			return responseSlice, err
		}

		// Add an element to the slice
		if counter == subcontainer.ShoppingId {

			container.Products = append(container.Products, ps)

		} else if counter == 0 {

			container.Products = append(container.Products, ps)

			counter = subcontainer.ShoppingId
			day = subcontainer.ShoppingDay

		} else {

			container = Response{ShoppingId: counter, ShoppingDay: day, Products: container.Products}
			responseSlice = append(responseSlice, container)

			container = Response{}

			container.Products = append(container.Products, ps)

			counter = subcontainer.ShoppingId
			day = subcontainer.ShoppingDay

			subcontainer = Response{}

		}

	}

	container = Response{ShoppingId: counter, ShoppingDay: day, Products: container.Products}
	responseSlice = append(responseSlice, container)

	return responseSlice, err
}
