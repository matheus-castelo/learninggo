package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/matheus-castelo/learninggo/src/configuration/rest_err/validation"
	"github.com/matheus-castelo/learninggo/src/controller/model/request"
	"github.com/matheus-castelo/learninggo/src/controller/model/response"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	log.Println("Init CreateUser controller")
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("Error trying to marshal object, error=%s\n", err.Error())
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	fmt.Println(userRequest)
	
	res := response.UserResponse{
		ID:    "test",
		Email: userRequest.Email,
	}

	c.JSON(http.StatusOK, res)
}