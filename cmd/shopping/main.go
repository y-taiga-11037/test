package main

import (
	db "gh.iiji.jp/y-taiga/mdtd_bootcamp/internal/database"
	"gh.iiji.jp/y-taiga/mdtd_bootcamp/internal/router"
	"gh.iiji.jp/y-taiga/mdtd_bootcamp/internal/server"
)

func main() {

	db.DB = db.Connect()
	defer db.DB.Close()
	router.Routing()
	server.Server()

}
