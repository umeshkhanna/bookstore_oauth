package app

import (
	"bookstore_oauth/src/domain/access_token"
	"bookstore_oauth/src/http"
	"bookstore_oauth/src/repository/db"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atService := access_token.NewService(db.NewRepository())
	atHandler := http.NewHandler(atService)
	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.Run(":8080")
}
