package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/otaviomart1ns/finsys/db/sqlc"
)

type Server struct {
	store  *db.SQLStore
	router *gin.Engine
}

/* func CORSConfig() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		context.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		context.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		context.Writer.Header().Set("Access-Control-Allow-Methods", "POST, DELETE, GET, PUT")

		if context.Request.Method == "OPTIONS" {
			context.AbortWithStatus(204)
			return
		}

		context.Next()
	}
} */

func NewServer(store *db.SQLStore) *Server {
	server := &Server{store: store}
	router := gin.Default()
	//router.Use(CORSConfig())

	router.POST("/user", server.addUser)
	router.GET("/user", server.getUsers)
	router.GET("/user/get-id/:id", server.getUserByID)
	router.GET("/user/get-email/:email", server.getUserByEmail)
	router.GET("/user/get-username/:username", server.getUserByUsername)
	router.GET("/user/get-name-lastname/:name/:last_name", server.getUserByNameAndLastName)
	router.GET("/user/get-email-password/:email/:password", server.getUserByEmailAndPassword)

	server.router = router
	return server

}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
