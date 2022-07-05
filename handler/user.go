package handler

import (
	"bank/auth"
	"bank/helper"
	"bank/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

// CreateUser godoc
// @Summary      Create User
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        body     body     user.RegisterUserInput     false  "example body"   user.RegisterUserInput
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
	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil {
		response := helper.APIResponse("cannot create token", http.StatusUnprocessableEntity, "error", err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		helper.NewError(c, http.StatusBadRequest, err)
		return
	}
	formatted := user.FormatUser(newUser, token)
	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatted)
	c.JSON(http.StatusOK, response)

}

// LoginUser godoc
// @Summary      Login User
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        body     body     user.LoginInput     false  "example body"   user.LoginInput
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
	token, err := h.authService.GenerateToken(loginUser.ID)
	if err != nil {
		response := helper.APIResponse("cannot create token", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		helper.NewError(c, http.StatusBadRequest, err)
		return
	}
	formatted := user.FormatUser(loginUser, token)
	response := helper.APIResponse("Login Success", http.StatusOK, "success", formatted)
	c.JSON(http.StatusOK, response)

}

// CheckEmail godoc
// @Summary      Check Email Availibility
// @Tags         email
// @Accept       json
// @Produce      json
// @Param        body     body     user.CheckEmailInput     false  "example body"   user.CheckEmailInput
// @Success      200  {object}  user.User
// @Failure      400  {object}  helper.HTTPError
// @Failure      404  {object}  helper.HTTPError
// @Failure      500  {object}  helper.HTTPError
// @Router       /email_checkers [post]
func (h *userHandler) CheckEmail(c *gin.Context) {
	var input user.CheckEmailInput
	//binding input to json
	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helper.APIResponse("Check Email Failed", http.StatusUnprocessableEntity, "error", err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		helper.NewError(c, http.StatusUnprocessableEntity, err)
		return
	}

	isEmailAvailable, err := h.userService.CheckEmail(input)
	if err != nil {
		response := helper.APIResponse("Email not available", http.StatusInternalServerError, "error", err.Error())
		c.JSON(http.StatusUnavailableForLegalReasons, response)
		helper.NewError(c, http.StatusUnprocessableEntity, err)
		return
	}
	// data := gin.H{
	// 	"is_available": isEmailAvailable
	// }
	var message string
	if isEmailAvailable {
		message = "Email is Available"
	} else {
		message = "Email has been registered"
	}
	formatter := user.FormatCheckEmail(isEmailAvailable)
	response := helper.APIResponse(message, http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}

// UploadAvatar godoc
// @Summary      Upload Avatar of user
// @Tags         email
// @Accept       json
// @Produce      json
// @Param        formData     formData     user.UploadAvatarInput     false  "example body"   collectionFormat(multi)
// @Success      200  {object}  user.User
// @Failure      400  {object}  helper.HTTPError
// @Failure      404  {object}  helper.HTTPError
// @Failure      500  {object}  helper.HTTPError
// @Router       /upload_avatar [post]
func (h *userHandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		formatter := user.FormatUploadAvatar(false)
		response := helper.APIResponse("Failed upload avatar", http.StatusBadRequest, "error", formatter)
		c.JSON(http.StatusBadRequest, response)
	}
	//make path of file
	currenUserLogin := c.MustGet("currenUserLogin")
	mapToUser := currenUserLogin.(user.User)
	userID := mapToUser.ID
	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)
	// path := "images/" + file.Filename
	//save file ke local
	err = c.SaveUploadedFile(file, path)

	if err != nil {
		formatter := user.FormatUploadAvatar(false)
		response := helper.APIResponse("Failed upload avatar", http.StatusBadRequest, "error", formatter)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.UploadAvatar(userID, path)

	if err != nil {
		formatter := user.FormatUploadAvatar(false)
		response := helper.APIResponse("Failed upload avatar", http.StatusBadRequest, "error", formatter)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUploadAvatar(true)
	response := helper.APIResponse("Successfully upload avatar", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}
