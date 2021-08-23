package main

import (
	"gh.iiji.jp/y-taiga/mdtd_bootcamp/internal/router"
	"gh.iiji.jp/y-taiga/mdtd_bootcamp/internal/server"
)

func main() {

	router.Routing()
	server.Server()
}
