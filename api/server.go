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

	router.POST("/users", server.addUser)
	router.PUT("/users", server.updateUser)
	router.DELETE("/users/:id", server.deleteUser)
	router.GET("/users", server.getUsers)
	router.GET("/users/:id", server.getUserByID)
	router.GET("/users/email/:email", server.getUserByEmail)
	router.GET("/users/username/:username", server.getUserByUsername)
	router.GET("/users/name/:name/lastname/:last_name", server.getUserByNameAndLastName)
	router.GET("/user/email/:email/password/:password", server.getUserByEmailAndPassword)

	router.POST("/categories", server.addCategory)
	router.PUT("/categories", server.updateCategory)
	router.DELETE("/categories/:id", server.deleteCategory)
	router.GET("/categories", server.getCategories) //rever rota, nao esta funcionando
	router.GET("/categories/:id", server.getCategoryByID)

	router.POST("/accounts", server.addAccount)
	router.PUT("/accounts", server.updateAccount)
	router.DELETE("/accounts/:id", server.deleteAccount)
	router.GET("/accounts", server.getAccounts) //rever rota, nao esta funcionando
	router.GET("/accounts/:id", server.getAccountByID)
	router.GET("/user/:user_id/accounts", server.getAccountByUser)
	router.GET("/category/:category_id/accounts", server.getAccountByCategory)
	router.GET("/accounts/graph/:graph", server.getAccountGraph)
	router.GET("/accounts/reports/:reports", server.getAccountReports)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
