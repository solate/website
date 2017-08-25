package controllers

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/solate/website/models"
)

//直接跳转到首页
func Home(c echo.Context) error {
	err := c.Redirect(http.StatusTemporaryRedirect, "/index.html")
	return err
}



func HomeShow(c echo.Context) (err error) {

	totalCount, err := models.GetStencilCount()
	models.CheckError(err)

	//导航列表
	navigationBarList, err := models.GetNavigationBarList()
	models.CheckError(err)

	//轮播图列表
	carouselList, err := models.GetCarouselList()
	models.CheckError(err)


	//模板列表
	stencilList, err := models.GetStencilList()
	models.CheckError(err)


	reusltDetail := make(map[string]interface{})
	reusltDetail["stencilList"] = stencilList
	reusltDetail["navigationBarList"] = navigationBarList
	reusltDetail["carouselList"] = carouselList


	c.JSON(http.StatusOK, models.ExportJson(models.Right_Status, "首页数据", reusltDetail, totalCount))
	return
}