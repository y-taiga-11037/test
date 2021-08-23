package database

import (
	"encoding/json"

	_ "github.com/go-sql-driver/mysql"
)

var DB = Connect()

func InsertProduct(body []byte, lastInsertID int64) error {

	var secondInsert Response
	err := json.Unmarshal(body, &secondInsert)
	if err != nil {
		return err
	}

	insertStatementForProduct := "INSERT INTO product (shopping_id, product_name, price, quantity) VALUES (?, ?, ?, ?);"
	stmtForProduct, err := DB.Prepare(insertStatementForProduct)
	if err != nil {
		return err
	}

	for _, v := range secondInsert.Products {

		_, err := stmtForProduct.Exec(lastInsertID, &v.ProductName, &v.Price, &v.Quantity)
		if err != nil {
			return err
		}
	}

	stmtForProduct.Close()
	return err
}
