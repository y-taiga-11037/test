package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	db "gh.iiji.jp/y-taiga/mdtd_bootcamp/internal/database"

	_ "github.com/go-sql-driver/mysql"
)

func GetShoppingListsHandler(w http.ResponseWriter, r *http.Request) {

	responseSlice, err := db.GetSelectShoppingListsQuery()
	if err != nil {
		log.Fatal(err)
		return
	}

	// Encode to json
	res, err := json.Marshal(responseSlice)
	if err != nil {
		log.Fatal(err)
		return
	}
	// Return json as a response
	w.Write(res)
}

func PostShoppingListsHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}

	lastInsertID, err := db.InsertShopping(body)
	if err != nil {
		log.Fatal(err)
	}

	err = db.InsertProduct(body, lastInsertID)
	if err != nil {
		log.Fatal(err)
	}

	responseSlice, err := db.GetPostQuery(lastInsertID)
	if err != nil {
		log.Fatal(err)
	}

	// Encode to json
	res, err := json.Marshal(responseSlice)
	if err != nil {
		log.Fatal(err)
	}
	// Return json as a response
	w.Write(res)
}
