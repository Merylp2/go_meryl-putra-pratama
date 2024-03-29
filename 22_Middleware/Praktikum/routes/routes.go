package routes

import (
	"test/22_Middleware/Praktikum/constants"
	"test/22_Middleware/Praktikum/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	jwtMiddleware := middleware.JWT([]byte(constants.SECRET_JWT))
	// Route / to handler function
	e.GET("/users", controllers.GetUsersController, jwtMiddleware)
	e.GET("/users/:id", controllers.GetUserController, jwtMiddleware)
	e.POST("/users", controllers.CreateUserController)
	e.POST("/users/login", controllers.LoginUserController)
	e.DELETE("/users/:id", controllers.DeleteUserController, jwtMiddleware)
	e.PUT("/users/:id", controllers.UpdateUserController, jwtMiddleware)
	e.GET("/books", controllers.GetBooksController, jwtMiddleware)
	e.GET("/books/:id", controllers.GetBookController, jwtMiddleware)
	e.POST("/books", controllers.CreateBookController, jwtMiddleware)
	e.DELETE("/books/:id", controllers.DeleteBookController, jwtMiddleware)
	e.PUT("/books/:id", controllers.UpdateBookController, jwtMiddleware)

	return e
}
