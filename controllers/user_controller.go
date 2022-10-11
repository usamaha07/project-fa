package controllers

import (
	"net/http"
	"project-fa/helpers"
	"project-fa/models"
	"project-fa/services"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userService services.UserServiceInterface
}

func NewUserController(userService services.UserServiceInterface) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) CreateUser(c echo.Context) error {
	var newUser models.CreateUserRequest
	err := c.Bind(&newUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.APIResponseFailed(err.Error()))
	}

	ctx := c.Request().Context()
	errCreateUser := uc.userService.CreateUser(ctx, newUser)
	if errCreateUser != nil {
		return c.JSON(http.StatusInternalServerError, helpers.APIResponseFailed(errCreateUser.Error()))
	}

	return c.JSON(http.StatusOK, helpers.APIResponseSuccessWithoutData("success create user"))
}
