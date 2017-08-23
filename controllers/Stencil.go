package controllers

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/solate/website/models"
	"strconv"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"time"
	"strings"
)

//模版展现
func StencilShow(c echo.Context) (err error) {

	pageStr := c.Param("page")
	stepStr := c.Param("step")


	page, err := strconv.Atoi(pageStr)
	models.CheckError(err)

	step, err := strconv.Atoi(stepStr)
	models.CheckError(err)

	reusltDetail, totalCount := showStencil(page, step)

	c.JSON(http.StatusOK, models.ExportJson(models.Right_Status, "模板列表", reusltDetail, totalCount))
	return
}
//显示
func showStencil(page, step int) (reusltDetail map[string]interface{}, totalCount int) {
	totalCount, err := models.GetStencilCount()
	models.CheckError(err)

	dataList, err := models.GetAllStencil(page, step)
	models.CheckError(err)

	reusltDetail = make(map[string]interface{})
	reusltDetail["dataList"] = dataList
	return
}

//模板添加
func StencilAdd(c echo.Context) (err error)  {

	//获得提交的json数据
	requestbody, err := ioutil.ReadAll(c.Request().Body)
	models.CheckError(err)

	type CurrentData struct {
		models.Stencil
		Step int `json:"step"`
		Page int `json:"page"`
	}

	var currentData CurrentData
	err = json.Unmarshal(requestbody, &currentData)
	models.CheckError(err)


	fmt.Println(currentData, "========1111==========")
	fmt.Println(currentData.Stencil, "========2222==========")


	//获得
	stencil := currentData.Stencil
	if stencil.Id > 0 {
		err = models.UpdateStencil(&stencil)
		models.CheckError(err)
	} else {
		//新增就获得新的id
		stencil.Id = models.GetId(models.DBStencil)
		stencil.Time = time.Now().Unix() //创建时间

		fmt.Println(currentData.Stencil, "========3333==========")

		err = models.AddStencil(&stencil)
		models.CheckError(err)

	}

	reusltDetail, totalCount := showStencil(currentData.Page, currentData.Step)

	c.JSON(http.StatusOK, models.ExportJson(models.Right_Status, "模板列表", reusltDetail, totalCount))
	return
}

//删除
func StencilDelete(c echo.Context) (err error)  {
	id := strings.TrimSpace(c.Param("id"))

	err = models.DeleteStencil(id)
	models.CheckError(err)

	c.JSON(http.StatusOK, models.ExportJson(models.Right_Status, nil, nil, -1))
	return
}