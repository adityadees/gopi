package controllers

import (
	"net/http"
	"time"
	"gopi/app/helper/utils"
	"gopi/config/constants"
	"github.com/labstack/echo"
	"github.com/dgrijalva/jwt-go"
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
		username,
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

func (ExampleController ExampleController) TestAPI(c echo.Context) error {

	return utils.HandleSuccessGet(c, constants.RESPONSE_SUCCESS_API_TEST)
}

func (ExampleController ExampleController) Accessible(c echo.Context) error {
	//return utils.HandleSuccessGet(c, constants.RESPONSE_SUCCESS_API_TEST)
	return c.String(http.StatusOK, "Accessible")
}