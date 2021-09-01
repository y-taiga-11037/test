package database

import (
	"encoding/json"

	_ "gh.iiji.jp/y-taiga/mdtd_bootcamp/internal/logger"
	_ "github.com/go-sql-driver/mysql"
	logging "github.com/sirupsen/logrus"
)

var DB = Connect()

func (c *ConnectDB) InsertProduct(body []byte, lastInsertID int64) error {

	var secondInsert Response
	err := json.Unmarshal(body, &secondInsert)
	if err != nil {
		logging.Error("Encoding failed")
		return err
	}
	logging.Info("Encoding completed")

	insertStatementForProduct := "INSERT INTO product (shopping_id, product_name, price, quantity) VALUES (?, ?, ?, ?);"
	stmtForProduct, err := DB.Prepare(insertStatementForProduct)
	if err != nil {
		logging.Error("Can't read Insert's text")
		return err
	}
	logging.Info("Read Insert's text")

	for _, v := range secondInsert.Products {

		_, err := stmtForProduct.Exec(lastInsertID, &v.ProductName, &v.Price, &v.Quantity)
		if err != nil {
			logging.Error("Insert failed")
			return err
		}
	}

	stmtForProduct.Close()
	logging.Info("Successfully inserted")
	return err
}
