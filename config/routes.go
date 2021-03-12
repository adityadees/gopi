package routes

import (
	"gopi/app/controllers"

	"github.com/labstack/echo"
)

type Routing struct {
	routes      controller.UserController
}