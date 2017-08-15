package controllers

import (
	"github.com/labstack/echo"
	"net/http"

)

//直接跳转到首页
func ManageIndex(c echo.Context) error {
	err := c.Redirect(http.StatusTemporaryRedirect, "/index.html")
	return err
}

