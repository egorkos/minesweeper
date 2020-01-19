package main

import "github.com/egorkos/minesweeper/app/interface/server"

func main() {
	router := server.CreateServer()
	router.Run()
}
