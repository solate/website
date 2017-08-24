package controllers

import (
	"github.com/labstack/echo"
	"strconv"
	"net/http"
	"encoding/json"
	"time"
	"strings"
	"io/ioutil"
	"github.com/solate/website/models"
	"fmt"
)

//导航栏展现
func NavigationBarShow(c echo.Context) (err error) {

	pageStr := c.Param("page")
	stepStr := c.Param("step")


	page, err := strconv.Atoi(pageStr)
	models.CheckError(err)

	step, err := strconv.Atoi(stepStr)
	models.CheckError(err)

	reusltDetail, totalCount := showNavigationBar(page, step)

	c.JSON(http.StatusOK, models.ExportJson(models.Right_Status, "导航列表", reusltDetail, totalCount))
	return
}
//显示
func showNavigationBar(page, step int) (reusltDetail map[string]interface{}, totalCount int) {
	totalCount, err := models.GetNavigationBarCount()
	models.CheckError(err)

	dataList, err := models.GetAllNavigationBar(page, step)
	models.CheckError(err)

	reusltDetail = make(map[string]interface{})
	reusltDetail["dataList"] = dataList
	return
}

//导航栏添加
func NavigationBarAdd(c echo.Context) (err error)  {

	//获得提交的json数据
	requestbody, err := ioutil.ReadAll(c.Request().Body)
	models.CheckError(err)

	type CurrentData struct {
		models.NavigationBar
		Step int `json:"step"`
		Page int `json:"page"`
	}

	var currentData CurrentData
	err = json.Unmarshal(requestbody, &currentData)
	models.CheckError(err)

	//获得
	navigationBar := currentData.NavigationBar
	if navigationBar.Id > 0 {
		err = models.UpdateNavigationBar(&navigationBar)
		models.CheckError(err)
	} else {
		//新增就获得新的id
		ids, err := models.GetId(models.DBNavigationBar)
		models.CheckError(err)

		navigationBar.Id = ids.Id
		navigationBar.Time = time.Now().Unix() //创建时间

		//添加
		err = models.AddNavigationBar(&navigationBar)
		models.CheckError(err)

	}

	reusltDetail, totalCount := showStencil(currentData.Page, currentData.Step)

	c.JSON(http.StatusOK, models.ExportJson(models.Right_Status, "模板列表", reusltDetail, totalCount))
	return
}

//导航栏删除
func NavigationBarDelete(c echo.Context) (err error)  {
	idStr := strings.TrimSpace(c.Param("id"))

	id, err := strconv.Atoi(idStr)
	models.CheckError(err)

	err = models.DeleteNavigationBar(id)
	models.CheckError(err)

	c.JSON(http.StatusOK, models.ExportJson(models.Right_Status, nil, nil, -1))
	return
}

func NavigationBarSearch(c echo.Context) (err error)  {

	name := c.Param("name")

	dataList, err := models.SearchNavigationBar(name)
	models.CheckError(err)

	reusltDetail := make(map[string]interface{})
	reusltDetail["dataList"] = dataList

	c.JSON(http.StatusOK, models.ExportJson(models.Right_Status, "模板列表", reusltDetail, -1))
	return
}