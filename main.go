package main

import (
	"log"
	"project-fa/controllers"
	"project-fa/databases"
	"project-fa/repositories"
	"project-fa/routes"
	"project-fa/services"

	"github.com/labstack/echo/v4"
)

func main() {
	// get connection to database mysql
	db := databases.GetConnectionMysql()
	defer db.Close()

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	authRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authController := controllers.NewAuthController(authService)

	e := echo.New()
	routes.UserPath(e, userController)
	routes.LoginAuth(e, authController)

	// starting web server
	log.Fatal(e.Start(":8080"))
}
