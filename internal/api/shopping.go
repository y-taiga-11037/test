package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

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

func (s *ShoppingHandler) CreateShoppingListsHandler(w http.ResponseWriter, r *http.Request) {

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

func (s *ShoppingHandler) UpdateDateHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "PATCH" {
		w.WriteHeader(http.StatusBadRequest)
		logging.Warning("A method that is not PATCH is specified")
		w.Write([]byte("A method that is not PATCH is specified"))
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		logging.Warning("Content-Type is not application/json")
		w.Write([]byte("Content-Type is not application/json"))
		return
	}

	logging.Infof("API request. method: %v, path: %v", r.Method, r.URL.Path)

	vars := mux.Vars(r)
	ID, err := strconv.Atoi(vars["shopping_id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logging.Error("Number conversion failure")
		w.Write([]byte("Number conversion failure"))
		return
	}
	var response db.DateResponse

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logging.Error("Error reading body data")
		w.Write([]byte("Error reading body data"))
		return
	}
	logging.Info("Body data loaded")

	response, err = s.shopping.UpdateDate(body, ID)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		logging.Error("Processing failed")
		w.Write([]byte("Processing failed"))
		return
	}
	logging.Info("Successful processing")

	// Encode to json
	res, err := json.Marshal(response)
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
	logging.Info("Date change process completed")
}

func (s *ShoppingHandler) DeleteShoppingListsHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "DELETE" {
		w.WriteHeader(http.StatusBadRequest)
		logging.Warning("A method that is not DELETE is specified")
		w.Write([]byte("A method that is not DELETE is specified"))
		return
	}

	logging.Infof("API request. method: %v, path: %v", r.Method, r.URL.Path)

	vars := mux.Vars(r)
	ID, err := strconv.Atoi(vars["shopping_id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logging.Error("Number conversion failure")
		w.Write([]byte("Number conversion failure"))
		return
	}

	err = s.shopping.DeleteProductByShoppingId(ID)
	if err != nil {
		logging.Error("Failed to delete ShoppingLists")
		return
	}

	err = s.shopping.DeleteShoppingByShoppingId(ID)
	if err != nil {
		logging.Error("Failed to delete ShoppingLists")
		return
	}
	logging.Info("ShoppingLists has been deleted successfully")
}

func (s *ShoppingHandler) GetSingleShoppingListHandler(w http.ResponseWriter, r *http.Request) {

	logging.Infof("API request. method: %v, path: %v", r.Method, r.URL.Path)

	vars := mux.Vars(r)
	ID, err := strconv.Atoi(vars["shopping_id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logging.Error("Number conversion failure")
		w.Write([]byte("Number conversion failure"))
		return
	}

	response, err := s.shopping.GetSingleShoppingList(ID)
	if err != nil {
		logging.Error("Failed to retrieve ShoppingLists")
		return
	}
	logging.Info("Successfully retrieved ShoppingLists")

	// Encode to json
	res, err := json.Marshal(response)
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

func (s *ShoppingHandler) UpdateShoppingListHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "PATCH" {
		w.WriteHeader(http.StatusBadRequest)
		logging.Warning("A method that is not PATCH is specified")
		w.Write([]byte("A method that is not PATCH is specified"))
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		logging.Warning("Content-Type is not application/json")
		w.Write([]byte("Content-Type is not application/json"))
		return
	}

	logging.Infof("API request. method: %v, path: %v", r.Method, r.URL.Path)

	vars := mux.Vars(r)
	ID, err := strconv.Atoi(vars["shopping_id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logging.Error("Number conversion failure")
		w.Write([]byte("Number conversion failure"))
		return
	}

	ProductID, err := strconv.Atoi(vars["shopping_product_id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logging.Error("Number conversion failure")
		w.Write([]byte("Number conversion failure"))
		return
	}

	var response db.FlagResponse

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logging.Error("Error reading body data")
		w.Write([]byte("Error reading body data"))
		return
	}
	logging.Info("Body data loaded")

	response, err = s.shopping.ChangeShoppingList(body, ID, ProductID)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		logging.Error("Processing failed")
		w.Write([]byte("Processing failed"))
		return
	}
	logging.Info("Successful processing")

	// Encode to json
	res, err := json.Marshal(response)
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
	logging.Info("Date change process completed")

}

func (s *ShoppingHandler) DeleteItemFromShoppingListHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "DELETE" {
		w.WriteHeader(http.StatusBadRequest)
		logging.Warning("A method that is not DELETE is specified")
		w.Write([]byte("A method that is not DELETE is specified"))
		return
	}

	logging.Infof("API request. method: %v, path: %v", r.Method, r.URL.Path)

	vars := mux.Vars(r)
	ID, err := strconv.Atoi(vars["shopping_id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logging.Error("Number conversion failure")
		w.Write([]byte("Number conversion failure"))
		return
	}

	ProductID, err := strconv.Atoi(vars["shopping_product_id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logging.Error("Number conversion failure")
		w.Write([]byte("Number conversion failure"))
		return
	}

	err = s.shopping.DeleteProduct(ID, ProductID)
	if err != nil {
		logging.Error("Failed to delete Product")
		return
	}

	logging.Info("Product has been deleted successfully")

}

func (s *ShoppingHandler) CreateItemFromShoppingListHandler(w http.ResponseWriter, r *http.Request) {

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

	vars := mux.Vars(r)
	ID, err := strconv.ParseInt(vars["shopping_id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logging.Error("Number conversion failure")
		w.Write([]byte("Number conversion failure"))
		return
	}

	err = s.shopping.InsertProduct(body, ID)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		logging.Error("Insert failed for Product table")
		w.Write([]byte("Insert failed for Product table"))
		return
	}
	logging.Info("Successfully inserted Product table")

	responseSlice, err := s.shopping.GetInsertLists(ID)
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
