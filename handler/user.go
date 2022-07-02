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

// CreateUser godoc
// @Summary      Create User
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        body     body     user.RegisterUserInput     false  " example"   user.RegisterUserInput
// @Success      200  {object}  helper.Response
// @Failure      400  {object}  helper.HTTPError
// @Failure      404  {object}  helper.HTTPError
// @Failure      500  {object}  helper.HTTPError
// @Router       /users [post]  CreateUser
func (h *userHandler) RegisterUser(c *gin.Context) {
	//take input from user
	var input user.RegisterUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		helper.NewError(c, http.StatusBadRequest, err)
		return
	}
	newUser, err := h.userService.RegisterUserInput(input)
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		helper.NewError(c, http.StatusBadRequest, err)
		return
	}
	// token,err := h.jwtService.GenerateToken()
	formatted := user.FormatUser(newUser, "token")
	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatted)
	c.JSON(http.StatusOK, response)

}

// LoginUser godoc
// @Summary      Login User
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        body     body     user.LoginInput     false  " example"   user.LoginInput
// @Success      200  {object}  user.User
// @Failure      400  {object}  helper.HTTPError
// @Failure      404  {object}  helper.HTTPError
// @Failure      500  {object}  helper.HTTPError
// @Router       /login [post]
func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput
	//binding input to json
	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		helper.NewError(c, http.StatusUnprocessableEntity, err)
		return
	}
	loginUser, err := h.userService.Login(input)
	formatted := user.FormatUser(loginUser, "token")
	response := helper.APIResponse("Login Success", http.StatusOK, "success", formatted)
	c.JSON(http.StatusOK, response)

}
