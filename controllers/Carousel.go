package controllers

import (
	"github.com/labstack/echo"
	"strconv"
	"net/http"
	"encoding/json"
	"time"
	"strings"
	"github.com/solate/website/models"
	"io/ioutil"
)

//模版展现
func CarouselShow(c echo.Context) (err error) {

	pageStr := c.Param("page")
	stepStr := c.Param("step")


	page, err := strconv.Atoi(pageStr)
	models.CheckError(err)

	step, err := strconv.Atoi(stepStr)
	models.CheckError(err)

	reusltDetail, totalCount := showCarousel(page, step)

	c.JSON(http.StatusOK, models.ExportJson(models.Right_Status, "轮播图列表", reusltDetail, totalCount))
	return
}
//显示
func showCarousel(page, step int) (reusltDetail map[string]interface{}, totalCount int) {
	totalCount, err := models.GetCarouselCount()
	models.CheckError(err)

	dataList, err := models.GetAllCarousel(page, step)
	models.CheckError(err)

	reusltDetail = make(map[string]interface{})
	reusltDetail["dataList"] = dataList
	return
}

//模板添加
func CarouselAdd(c echo.Context) (err error)  {

	//获得提交的json数据
	requestbody, err := ioutil.ReadAll(c.Request().Body)
	models.CheckError(err)

	type CurrentData struct {
		models.Carousel
		Step int `json:"step"`
		Page int `json:"page"`
	}

	var currentData CurrentData
	err = json.Unmarshal(requestbody, &currentData)
	models.CheckError(err)

	//获得
	carousel := currentData.Carousel
	if carousel.Id > 0 {
		err = models.UpdateCarousel(&carousel)
		models.CheckError(err)
	} else {
		//新增就获得新的id
		ids, err := models.GetId(models.DBCarousel)
		models.CheckError(err)

		carousel.Id = ids.Id
		carousel.Time = time.Now().Unix() //创建时间

		//添加
		err = models.AddCarousel(&carousel)
		models.CheckError(err)

	}

	reusltDetail, totalCount := showCarousel(currentData.Page, currentData.Step)

	c.JSON(http.StatusOK, models.ExportJson(models.Right_Status, "模板列表", reusltDetail, totalCount))
	return
}

//删除
func CarouselDelete(c echo.Context) (err error)  {
	idStr := strings.TrimSpace(c.Param("id"))

	id, err := strconv.Atoi(idStr)
	models.CheckError(err)

	err = models.DeleteCarousel(id)
	models.CheckError(err)

	c.JSON(http.StatusOK, models.ExportJson(models.Right_Status, nil, nil, -1))
	return
}

func CarouselSearch(c echo.Context) (err error)  {

	//获得提交的json数据
	requestbody, err := ioutil.ReadAll(c.Request().Body)
	models.CheckError(err)

	type CurrentData struct {
		SearchImage string `json:"SearchImage"`
	}

	var currentData CurrentData
	err = json.Unmarshal(requestbody, &currentData)
	models.CheckError(err)

	dataList, err := models.SearchCarousel(currentData.SearchImage)
	models.CheckError(err)

	reusltDetail := make(map[string]interface{})
	reusltDetail["dataList"] = dataList

	c.JSON(http.StatusOK, models.ExportJson(models.Right_Status, "模板列表", reusltDetail, -1))
	return
}