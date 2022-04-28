package main

import (
	"github.com/jclaudiotomasjr/alura/apip-gin-rest-app/database"
	"github.com/jclaudiotomasjr/alura/apip-gin-rest-app/routes"
)

func main() {
	database.ConnectDatabase()
	routes.HandleRequests()
}
