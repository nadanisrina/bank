package handler

import (
	"bank/helper"
	"bank/user"
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
	//take input from user
	var input user.RegisterUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}
	user, err := h.userService.RegisterUserInput(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", user)
	c.JSON(http.StatusOK, response)

}
