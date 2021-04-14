package routes

import (
	"net/http"
	"gopi/app/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/dgrijalva/jwt-go"
)

type Routing struct {
	example      controllers.ExampleController
}

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}


func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	name := claims.Name
	return c.String(http.StatusOK, "Welcome "+name+"!")
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
