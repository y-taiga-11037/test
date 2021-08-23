package server

import (
	"log"
	"net/http"

	db "gh.iiji.jp/y-taiga/mdtd_bootcamp/internal/database"
	"gh.iiji.jp/y-taiga/mdtd_bootcamp/internal/router"
)

func Server() {

	log.Fatal(http.ListenAndServe(":8080", router.R))
	defer db.DB.Close()
}
