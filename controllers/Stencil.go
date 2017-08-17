package controllers

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/solate/website/models"
	"strconv"
	"fmt"
)

//模版展现
func StencilShow(c echo.Context) (err error) {

	pageStr := c.Param("page")
	stepStr := c.Param("step")


	page, err := strconv.Atoi(pageStr)
	models.CheckError(err)

	step, err := strconv.Atoi(stepStr)
	models.CheckError(err)

	totalCount, err := models.GetStencilCount()
	models.CheckError(err)

	dataList, err := models.GetAllStencil(page, step)
	models.CheckError(err)

	fmt.Println(dataList, "===========")

	reusltDetail := make(map[string]interface{})
	reusltDetail["dataList"] = dataList

	c.JSON(http.StatusOK, models.ExportJson(models.Right_Status, "模板列表", reusltDetail, totalCount))
	return
}
