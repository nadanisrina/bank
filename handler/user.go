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
// @Description  create user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      object  true  "user.RegisterUserInput"
// @Success      200  {object}  user.User
// @Failure      400  {object}  helper.HTTPError
// @Failure      404  {object}  helper.HTTPError
// @Failure      500  {object}  helper.HTTPError
// @Router       /users [post]
func (h *userHandler) RegisterUser(c *gin.Context) {
	//take input from user
	var input user.RegisterUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		helper.NewError(c, http.StatusBadRequest, err)
		return
	}
	newUser, err := h.userService.RegisterUserInput(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		helper.NewError(c, http.StatusBadRequest, err)
		return
	}
	// token,err := h.jwtService.GenerateToken()
	formatted := user.FormatUser(newUser, "token")
	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatted)
	c.JSON(http.StatusOK, response)

}
