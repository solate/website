package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

//登录页面
func Login(c echo.Context) error {
	err := c.Redirect(http.StatusFound, "/login.html")
	return err
}

//注册页面
func Register(c echo.Context) error {
	err := c.Redirect(http.StatusFound, "/register.html")
	return err
}

