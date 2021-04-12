package routes

import (
	"gopi/app/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Routing struct {
	example      controllers.ExampleController
}

func (Routing Routing) GetRoutes() *echo.Echo {
	e := echo.New()

	
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", Routing.example.Accessible)
	e.POST("/login", Routing.example.Login)
	r := e.Group("/restricted")
	config := middleware.JWTConfig{
		Claims:     &jwtCustomClaims{},
		SigningKey: []byte("mySecretKey"),
	}
	r.Use(middleware.JWTWithConfig(config))
	r.GET("", restricted)
	return e
}
