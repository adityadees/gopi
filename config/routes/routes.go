package routes

import (
	"gopi/app/controllers"
	"github.com/labstack/echo"
)

type Routing struct {
	userc      controllers.Userc
}

func (Routing Routing) GetRoutes() *echo.Echo {
	e := echo.New()
	e.POST("/login", Routing.userc.Login)
	e.POST("/", Routing.userc.xlog)

	return e
}
