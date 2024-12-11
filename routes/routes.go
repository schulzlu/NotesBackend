package routes

import (
	"github.com/gin-gonic/gin"
	"notes.com/app/middlewares"
)

func RegisterRoutes(server *gin.Engine) {

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/notes", CreateNote)
	authenticated.PUT("/notes/:id", UpdateNote)
	authenticated.DELETE("/notes/:id", DeleteNote)

	server.GET("/notes", GetNotes)
	server.GET("/notes/:id", GetSingleNote)

	server.POST("/login", LoginUser)
	server.GET("/users", GetUsers)
	server.GET("/users/:id", GetUserById)
	server.POST("/users", CreateUser)
}
