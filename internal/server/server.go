package server

import (
	"log"
	"net/http"

	_ "gh.iiji.jp/y-taiga/mdtd_bootcamp/internal/logger"
	"gh.iiji.jp/y-taiga/mdtd_bootcamp/internal/router"
	logging "github.com/sirupsen/logrus"
)

func Server() {

	logging.Info("Server Startup")
	log.Fatal(http.ListenAndServe(":8080", router.R))

}
