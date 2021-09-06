package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	db "gh.iiji.jp/y-taiga/mdtd_bootcamp/internal/database"
	_ "gh.iiji.jp/y-taiga/mdtd_bootcamp/internal/logger"
	logging "github.com/sirupsen/logrus"
)

type ShoppingHandler struct {
	shopping db.Database
}

func NewShoppingHandler(s db.Database) *ShoppingHandler {
	return &ShoppingHandler{s}
}

func (s *ShoppingHandler) GetShoppingListsHandler(w http.ResponseWriter, r *http.Request) {

	logging.Infof("API request. method: %v, path: %v", r.Method, r.URL.Path)

	responseSlice, err := s.shopping.GetShoppingLists()
	if err != nil {
		logging.Error("Failed to retrieve ShoppingLists")
		return
	}
	logging.Info("Successfully retrieved ShoppingLists")

	// Encode to json
	res, err := json.Marshal(responseSlice)
	if err != nil {
		logging.Error("Encoding failed")
		return
	}
	logging.Info("Encoding completed")

	// Return json as a response
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	logging.Info("GetShoppingLists process completed")
}

func (s *ShoppingHandler) PostShoppingListsHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		logging.Warning("A method that is not POST is specified")
		w.Write([]byte("A method that is not POST is specified"))
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		logging.Warning("Content-Type is not application/json")
		w.Write([]byte("Content-Type is not application/json"))
		return
	}

	logging.Infof("API request. method: %v, path: %v", r.Method, r.URL.Path)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logging.Error("Error reading body data")
		w.Write([]byte("Error reading body data"))
		return
	}
	logging.Info("Body data loaded")

	lastInsertID, err := s.shopping.InsertShopping(body)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		logging.Error("Insert failed for Shopping table")
		w.Write([]byte("Insert failed for Shopping table"))
		return
	}
	logging.Info("Successfully inserted Shopping table")

	err = s.shopping.InsertProduct(body, lastInsertID)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		logging.Error("Insert failed for Product table")
		w.Write([]byte("Insert failed for Product table"))
		return
	}
	logging.Info("Successfully inserted Product table")

	responseSlice, err := s.shopping.GetInsertLists(lastInsertID)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		logging.Error("Failed to retrieve ShoppingLists")
		w.Write([]byte("Failed to retrieve ShoppingLists"))
		return
	}
	logging.Info("Successfully retrieved ShoppingLists")

	// Encode to json
	res, err := json.Marshal(responseSlice)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		logging.Error("Encoding failed")
		w.Write([]byte("Encoding failed"))
		return
	}
	logging.Info("Encoding completed")

	// Return json as a response
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	logging.Info("PostShoppingLists process completed")
}
