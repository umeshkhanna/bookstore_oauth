package http

import (
	atDomain "bookstore_oauth/src/domain/access_token"
	"bookstore_oauth/src/services/access_token"
	"bookstore_oauth/src/utils/errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
	Create(*gin.Context)
	UpdateExpirationTime(*gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (handler *accessTokenHandler) GetById(c *gin.Context) {
	accessTokenId := strings.TrimSpace(c.Param("access_token_id"))
	accessToken, err := handler.service.GetById(accessTokenId)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}

func (handler *accessTokenHandler) Create(c *gin.Context) {
	var token atDomain.AccessTokenRequest
	if err := c.ShouldBindJSON(&token); err != nil {
		restErr := errors.NewBadRequestError("Invalid JSON Body")
		c.JSON(restErr.Status, restErr)
		return
	}
	accessToken, tokenErr := handler.service.Create(token)
	if tokenErr != nil {
		c.JSON(tokenErr.Status, tokenErr)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}

func (handler *accessTokenHandler) UpdateExpirationTime(c *gin.Context) {
	var token atDomain.AccessToken
	if err := c.ShouldBindJSON(&token); err != nil {
		restErr := errors.NewBadRequestError("Invalid JSON Body")
		c.JSON(restErr.Status, restErr)
		return
	}
	if err := handler.service.UpdateExpirationTime(token); err != nil {
		c.JSON(err.Status, err)
	}
	c.JSON(http.StatusOK, token)
}
