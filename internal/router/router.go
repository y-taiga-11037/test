package router

import (
	"github.com/gorilla/mux"

	api "gh.iiji.jp/y-taiga/mdtd_bootcamp/internal/api"
	_ "gh.iiji.jp/y-taiga/mdtd_bootcamp/internal/logger"
	logging "github.com/sirupsen/logrus"
)

var R = mux.NewRouter()

func Routing() {
	logging.Info("Router Startup")
	R.HandleFunc("/api/shopping", api.GetShoppingListsHandler).Methods("GET")
	R.HandleFunc("/api/shopping", api.PostShoppingListsHandler).Methods("POST")
}
