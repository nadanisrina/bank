package handler

import (
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
	c.JSON(http.StatusOK, user)
}
