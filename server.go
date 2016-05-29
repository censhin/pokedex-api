package main

import (
	"log"
	"net/http"
	"github.com/censhin/pokedex-api/db"
	"github.com/censhin/pokedex-api/utils"
)

var config = utils.GetConfig()
var port = config.Server.Port

func main() {
	db.Connect()
	defer db.Close()

	log.Println("Starting server on port " + port)
	http.ListenAndServe(":" + port, InitRoutes())
}
