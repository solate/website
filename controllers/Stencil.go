package controllers

import (
	"net/http"
	"io/ioutil"
	"github.com/labstack/echo"
	"github.com/Sirupsen/logrus"
	"encoding/json"
	"github.com/solate/website/models"
)

//模版展现
func StencilShow(c echo.Context) (err error) {

	requestbody, err := ioutil.ReadAll(c.Request().Body)
	models.CheckError(err)

	type receiveData struct {
		Page int `json:"page"`
		PageSize int `json:"pageSize"`
	}
	var currentData receiveData
	err = json.Unmarshal(requestbody, &currentData)
	if err != nil {
		logrus.Error(err)
		return err
	}

	totalCount, err := models.GetStencilCount()
	models.CheckError(err)

	dataList, err := models.GetAllStencil(currentData.Page, currentData.PageSize)
	models.CheckError(err)

	reusltDetail := make(map[string]interface{})
	reusltDetail["data"] = dataList

	c.JSON(http.StatusOK, models.ExportJson(models.Right_Status, "", reusltDetail, totalCount))
	return
}
