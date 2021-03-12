package controllers

import (
	"fmt"
	"net/http"
	"time"
	/*"gopi/app/handles"*/
	"gopi/app/models"
	/*"gopi/app/models/entities"*/
	"gopi/config"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Userc struct {
	model models.user_m
}
