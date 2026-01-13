package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheus-castelo/learninggo/src/configuration/rest_err"
	"github.com/matheus-castelo/learninggo/src/controller/model/request"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(fmt.Sprintf("There are some incorrect fields, error=%s", err.Error()))
		c.JSON(int(restErr.Code), restErr)
		return
	}

	fmt.Println(userRequest)
	c.JSON(http.StatusOK, userRequest)
}