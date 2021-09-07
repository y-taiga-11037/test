package database

import (
	"encoding/json"
	"fmt"

	_ "gh.iiji.jp/y-taiga/mdtd_bootcamp/internal/logger"
	_ "github.com/go-sql-driver/mysql"
	logging "github.com/sirupsen/logrus"
)

func (c *ConnectDB) InsertProduct(body []byte, lastInsertID int64) error {

	var secondInsert Response
	err := json.Unmarshal(body, &secondInsert)
	if err != nil {
		logging.Error("Encoding failed")
		return err
	}
	logging.Info("Encoding completed")

	fmt.Println(body)
	fmt.Println(secondInsert)

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

func (c *ConnectDB) DeleteProductTable(shopping_id int) error {

	_, err := DB.Exec("DELETE FROM product WHERE shopping_id = ?", shopping_id)
	if err != nil {
		logging.Error("ProductTable Deletion failure")
		return err
	}
	logging.Info("ProductTable Deletion Success")

	return err
}

func (c *ConnectDB) ChangeShoppingList(body []byte, shopping_id int, shopping_product_id int) (FlagResponse, error) {

	var update FlagResponse
	var patchcontents PurcheseFlag
	err := json.Unmarshal(body, &patchcontents)
	if err != nil {
		logging.Error("Encoding failed")
		return update, err
	}
	logging.Info("Encoding completed")

	_, err = DB.Exec("UPDATE product set product_name = ?, price = ?, quantity = ?, purchase_flag = ? WHERE shopping_id = ? AND shopping_product_id = ?", &patchcontents.ProductName, &patchcontents.Price, &patchcontents.Quantity, &patchcontents.PurcheseFlag, shopping_id, shopping_product_id)
	if err != nil {
		logging.Error("Couldn't update")
		return update, err
	}
	logging.Info("Update Complete")

	update = FlagResponse{ShoppingId: shopping_id, Products: patchcontents}

	return update, err

}

func (c *ConnectDB) DeleteProduct(shopping_id int, shopping_product_id int) error {

	_, err := DB.Exec("DELETE FROM product WHERE shopping_id = ? AND shopping_product_id = ?", shopping_id, shopping_product_id)
	if err != nil {
		logging.Error("Product Deletion failure")
		return err
	}
	logging.Info("Product Deletion Success")

	return err
}
