package router

import (
	"github.com/gorilla/mux"

	api "gh.iiji.jp/y-taiga/mdtd_bootcamp/internal/api"
	db "gh.iiji.jp/y-taiga/mdtd_bootcamp/internal/database"
	_ "gh.iiji.jp/y-taiga/mdtd_bootcamp/internal/logger"
	logging "github.com/sirupsen/logrus"
)

var R = mux.NewRouter()

func Routing() {
	logging.Info("Router Startup")

	h := api.NewGetHandler(&db.CategoryModelDB{})

	R.HandleFunc("/api/shopping", api.GetShoppingListsHandler).Methods("GET")
	R.HandleFunc("/api/shopping", api.PostShoppingListsHandler).Methods("POST")
}
