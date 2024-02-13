package main

import (
	"api_com_gin/database"
	"api_com_gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()

	routes.HandleRequest()
}
