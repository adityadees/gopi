package controllers

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	//"github.com/labstack/echo/v4"
	//"github.com/labstack/echo/v4/middleware"
	"gopi/app/helper/utils"
	"gopi/config/constants"

	"github.com/labstack/echo"
)

type ExampleController struct {
}

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

func (ExampleController ExampleController) Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Throws unauthorized error
	if username != "adityads" || password != "123" {
		return echo.ErrUnauthorized
	}

	// Set custom claims
	claims := &jwtCustomClaims{
		"Jon Snow",
		true,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("mySecretKey"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func (ExampleController ExampleController) Restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	name := claims.Name
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

func (ExampleController ExampleController) TestAPI(c echo.Context) error {

	return utils.HandleSuccessGet(c, constants.RESPONSE_SUCCESS_API_TEST)
}

func (ExampleController ExampleController) Accessible(c echo.Context) error {
	//return utils.HandleSuccessGet(c, constants.RESPONSE_SUCCESS_API_TEST)
	return c.String(http.StatusOK, "Accessible")
}