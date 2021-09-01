package database

import (
	"encoding/json"

	_ "gh.iiji.jp/y-taiga/mdtd_bootcamp/internal/logger"
	_ "github.com/go-sql-driver/mysql"
	logging "github.com/sirupsen/logrus"
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

type CategoryModelDB struct{}

func (c *CategoryModelDB) GetShoppingLists() ([]Response, error) {

	responseSlice := make([]Response, 0)

	// Execute a SELECT
	rows, err := DB.Query("SELECT P.shopping_id, S.shopping_day, P.product_name, P.quantity, P.price FROM shopping AS S LEFT OUTER JOIN product AS P ON P.shopping_id = S.shopping_id;")
	if err != nil {
		logging.Error("Can't read Insert's text")
		return responseSlice, err
	}
	logging.Info("Read Insert's text")

	var id int = 0
	var day string

	var container Response
	var subcontainer Response

	for rows.Next() {
		var ps Product

		// Map the SELECT result to a structure
		err = rows.Scan(&subcontainer.ShoppingId, &subcontainer.ShoppingDay, &ps.ProductName, &ps.Quantity, &ps.Price)
		if err != nil {
			logging.Error("Couldn't map the result of SELECT to a structure")
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

	logging.Info("Map the SELECT result to a structure")
	return responseSlice, err
}

func InsertShopping(body []byte) (int64, error) {

	var firstInsert Response
	err := json.Unmarshal(body, &firstInsert)
	if err != nil {
		logging.Error("Encoding failed")
		return 0, err
	}
	logging.Info("Encoding completed")

	insertStatement := "INSERT INTO shopping (shopping_day) VALUES (?);"
	stmt, err := DB.Prepare(insertStatement)
	if err != nil {
		logging.Error("Can't read Insert's text")
		return 0, err
	}
	logging.Info("Read Insert's text")

	result, err := stmt.Exec(&firstInsert.ShoppingDay)
	if err != nil {
		logging.Error("Unable to read Shoppingday value")
		return 0, err
	}
	logging.Info("Read Shoppingday value")

	stmt.Close()

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		logging.Error("Unable to get lastInsertID")
		return 0, err
	}
	logging.Info("Get lastInsertID")
	return lastInsertID, err
}

func GetInsertLists(lastInsertID int64) (Response, error) {

	var id int = 0
	var day string
	var container Response
	var subcontainer Response

	// Execute a SELECT
	rows, err := DB.Query("SELECT P.shopping_id, S.shopping_day, P.product_name, P.quantity, P.price FROM shopping AS S LEFT OUTER JOIN product AS P ON P.shopping_id = S.shopping_id WHERE P.shopping_id = CONCAT(?, '%')", lastInsertID)
	if err != nil {
		logging.Error("Can't read Insert's text")
		return container, err
	}
	logging.Info("Read Insert's text")

	for rows.Next() {
		var ps Product

		// Map the SELECT result to a structure
		err = rows.Scan(&subcontainer.ShoppingId, &subcontainer.ShoppingDay, &ps.ProductName, &ps.Quantity, &ps.Price)
		if err != nil {
			logging.Error("Couldn't map the result of SELECT to a structure")
			return container, err
		}

		// Add an element to the slice
		if id == subcontainer.ShoppingId {

			container.Products = append(container.Products, ps)

		} else if id == 0 {

			container.Products = append(container.Products, ps)

			id = subcontainer.ShoppingId
			day = subcontainer.ShoppingDay

		}

	}

	container = Response{ShoppingId: id, ShoppingDay: day, Products: container.Products}

	logging.Info("Map the SELECT result to a structure")
	return container, err
}
