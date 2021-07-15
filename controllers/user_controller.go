package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoangduyptithcm/bookstore_users_api/domain/users"
	"github.com/hoangduyptithcm/bookstore_users_api/services"
	"github.com/hoangduyptithcm/bookstore_users_api/utils/errors"
)

var (
	counter int
)

func CreateUser(c *gin.Context) {
	print("ccccccccccccccccc")
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		//todo: return bad request to the caller
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		//todo handle user create error
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusNotImplemented, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusOK, "GetUser")
}

func FindUserByName(c *gin.Context) {
	c.String(http.StatusOK, "FindUserByName")
}
