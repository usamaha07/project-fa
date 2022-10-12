package routes

import (
	"project-fa/controllers"
	"project-fa/middlerwares"

	"github.com/labstack/echo/v4"
)

func UserPath(e *echo.Echo, uc *controllers.UserController) {
	e.POST("/users", uc.CreateUser)
	e.GET("/users/:user_id", uc.GetUserById)
	e.GET("/users", uc.GetAllUser)
	e.DELETE("/users", uc.DeleteUser, middlerwares.JWTMiddleware())
}

func LoginAuth(e *echo.Echo, ac *controllers.AuthController) {
	e.POST("/login", ac.Login)
}
