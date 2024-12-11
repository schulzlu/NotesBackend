package main

import (
	"github.com/gin-gonic/gin"
	db "notes.com/app/database"
	"notes.com/app/routes"
)

func main() {
	db.InitDatabase()

	server := gin.Default()

	routes.RegisterRoutes(server)
	server.Run()
}
