package router

import (
	"github.com/gorilla/mux"

	api "gh.iiji.jp/y-taiga/mdtd_bootcamp/internal/api"
	db "gh.iiji.jp/y-taiga/mdtd_bootcamp/internal/database"
	logging "github.com/sirupsen/logrus"
)

var R = mux.NewRouter()

func Routing() {
	logging.Info("Router Startup")

	h := api.NewShoppingHandler(&db.ConnectDB{})

	R.HandleFunc("/api/shopping", h.GetShoppingListsHandler).Methods("GET")
	R.HandleFunc("/api/shopping", h.CreateShoppingListsHandler).Methods("POST")
	R.HandleFunc("/api/shopping/{shopping_id}", h.UpdateDateHandler).Methods("PATCH")
	R.HandleFunc("/api/shopping/{shopping_id}", h.DeleteShoppingListsHandler).Methods("DELETE")
	R.HandleFunc("/api/shopping/{shopping_id}", h.GetSingleShoppingListHandler).Methods("GET")
	R.HandleFunc("/api/shopping/{shopping_id}/products/{shopping_product_id}", h.UpdateShoppingListHandler).Methods("PATCH")
	R.HandleFunc("/api/shopping/{shopping_id}/products/{shopping_product_id}", h.DeleteItemFromShoppingListHandler).Methods("DELETE")
	R.HandleFunc("/api/shopping/{shopping_id}/products", h.CreateItemFromShoppingListHandler).Methods("POST")

}
