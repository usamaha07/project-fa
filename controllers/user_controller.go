package controllers

import (
	"net/http"
	"project-fa/helpers"
	"project-fa/middlerwares"
	"project-fa/models"
	"project-fa/services"
	"strconv"

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

func (uc *UserController) GetUserById(c echo.Context) error {

	idString := c.Param("user_id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.APIResponseFailed("id not recognise"))
	}

	ctx := c.Request().Context()
	user, err := uc.userService.GetUserById(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.APIResponseFailed(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.APIResponseSuccess("succes get user", user))
}

func (uc *UserController) GetAllUser(c echo.Context) error {
	ctx := c.Request().Context()
	users, err := uc.userService.GetAllUser(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.APIResponseFailed(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.APIResponseSuccess("success get all user", users))
}

func (uc *UserController) DeleteUser(c echo.Context) error {
	// get id user from token
	idToken, errToken := middlerwares.ExtractToken(c)
	if errToken != nil {
		return c.JSON(http.StatusUnauthorized, helpers.APIResponseFailed("unauthorized"))
	}

	ctx := c.Request().Context()
	err := uc.userService.DeleteUser(ctx, idToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.APIResponseFailed(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.APIResponseSuccessWithoutData("success delete user"))
}
