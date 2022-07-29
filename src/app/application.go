package app

import (
	"bookstore_oauth/src/http"
	"bookstore_oauth/src/repository/db"
	"bookstore_oauth/src/repository/rest"
	"bookstore_oauth/src/services/access_token"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atService := access_token.NewService(rest.NewRepository(), db.NewRepository())
	atHandler := http.NewHandler(atService)
	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)
	router.Run(":8081")
}
