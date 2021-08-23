package router

import (
	"github.com/gorilla/mux"

	api "gh.iiji.jp/y-taiga/mdtd_bootcamp/internal/api"
)

var R = mux.NewRouter()

func Routing() {

	R.HandleFunc("/api/shopping", api.GetShoppingListsHandler).Methods("GET")
	R.HandleFunc("/api/shopping", api.PostShoppingListsHandler).Methods("POST")
}
