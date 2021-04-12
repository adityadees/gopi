package utils

import (
	"net/http"
	"github.com/labstack/echo"
	"gopi/app/helper/utils/response"
	"gopi/config/constants"
)


func HandleSuccessGet(c echo.Context, data interface{}) error {
	res := response.ResponseGenericGet{
		Status:  constants.SUCCESS,
		Message: constants.SUCCESS_LOAD_DATA,
		Data:    data,
	}
	return c.JSON(http.StatusOK, res)
}

/*func HandleSuccessPost(c echo.Context, message string) error {
	res := response.ResponseGenericIn{
		Status: constants.SUCCESS,
		Message: message,
	}
	return c.JSON(http.StatusOK, res)
}

func HandleError(c echo.Context, status int, message string) error {
	res := response.ResponseGenericIn{
		Status:  constants.ERROR,
		Message: message,
	}
	return c.JSON(status, res)
}*/