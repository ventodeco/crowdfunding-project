package handler

import (
	"bwastartup/helper"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	// catch request from user
	// map request from user to RegisterUserRequest
	// struct at above need to passing as service parameter

	var request user.RegisterUserRequest

	err := c.ShouldBindJSON(&request)

	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	newUser, err := h.userService.RegisterUser(request)

	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	formatter := user.FormatUser(newUser, "token")

	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}
