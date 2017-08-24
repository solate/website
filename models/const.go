package models

import (
	"github.com/Sirupsen/logrus"
)

const (
	//mongo使用哪个数据库存储
	DB = "test"

	//自增id管理
	DBIds = "ids"
	//模块管理
	DBStencil = "stencil"
	//导航
	DBNavigationBar = "navigation"
	//轮播图
	DBCarousel = "carousel"


	//状态码
	Right_Status  = 200
	ACCESS_DENIED = 1005
	VALUE_EXISTS  = 2007
	NOT_FOUND     = 4001
	SERVER_ERROR  = 5001
)

var _errors = map[int]string{
	4001: "not found",
	5001: "server error",
}

func CheckError(err error) {
	if err != nil {
		logrus.Error(err)
	}
}

func ExportJson(reusltCode int, title, Content interface{}, totalCount int) map[string]interface{} {

	if reusltCode == Right_Status {
		//如果是正常的输出
		var resultMap map[string]interface{} = map[string]interface{}{
			"reusltCode":   Right_Status,
			"reusltTitle":  title,
			"reusltDetail": Content,
		}
		if totalCount >= 0 {
			resultMap["totalCount"] = totalCount
		}

		return resultMap
	} else {
		//如果是明确的错误输出
		reusltTitle := _errors[reusltCode]
		return map[string]interface{}{
			"reusltCode":   reusltCode,
			"reusltTitle":  reusltTitle,
			"reusltDetail": Content,
		}
	}
	//如果是异常的错误输出
	return map[string]interface{}{
		"reusltCode":   SERVER_ERROR,
		"reusltTitle":  "unknown error occured",
		"reusltDetail": "unknown error occured",
	}
}

