package controllers

import (
	"fmt"
	validation2 "github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"go-name/src/calendar"
	"go-name/src/models"
	"go-name/src/pkg/app"
	"go-name/src/pkg/e"
	"go-name/src/pkg/logging"
	"net/http"
)

type NameRequest struct {
	Year     int `json:"year"`
	Month    int `json:"month"`
	Day      int `json:"day"`
	Hour     int `json:"hour"`
	SurName  int `json:"surName"`
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
}

type NameResponse struct {
	*models.TName
	*calendar.SolarDate
}

// GetNames 获取用户名字分析
func GetNames(c *gin.Context) {
	appG := app.Gin{C: c}
	var body NameRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		appG.Response(http.StatusBadRequest, e.InvalidParams, err.Error())
		return
	}

	logging.Info("GetNames body: ", body)
	valid := validation2.Validation{}

	valid.Range(body.Year, 1950, 2100, "year")
	valid.Range(body.Month, 1, 12, "month")
	valid.Range(body.Day, 1, 31, "day")
	valid.Min(body.PageNum, 0, "pageNumber")
	valid.Range(body.PageSize, 1, 30, "pageSize")

	if valid.HasErrors() {
		appG.Response(http.StatusBadRequest, e.InvalidParams, valid.Errors)
		return
	}
	birthTime := calendar.DateReturn{
		Year:  body.Year,
		Month: body.Month,
		Day:   body.Day,
		Hour:  body.Hour,
	}.GetBirthTime()

	solarDate := calendar.GetSolarDate(birthTime)

	query := make(map[string]interface{})

	//query["five_elements"] = solarDate.GanZhi.PositiveGod
	names, err := models.GetNames(body.PageNum, body.PageSize, query)

	if err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR_QUERY_DB_FAIL, err.Error())
		return
	}

	data := make(map[string]interface{})

	if len(names) == 0 {

		appG.Response(http.StatusOK, e.SUCCESS, data)
		return
	}

	var wordQueryArray []string
	for _, nameVal := range names {
		name := nameVal.Name
		for _, word := range name {
			wordQueryArray = append(wordQueryArray, string(word))
		}
	}

	fmt.Printf("%#v\n", wordQueryArray)
	wordsResult, err := models.GetWords(wordQueryArray)

	// 不报错，记录错误日志
	if err != nil {
		logging.Info(err.Error())
		//appG.Response(http.StatusBadRequest, e.ERROR_QUERY_DB_FAIL, error.Error())
		//return
	}

	var nameResponse []NameResponse

	for _, value := range names {
		//组装八字、姓名
		name := NameResponse{value, solarDate}
		nameResponse = append(nameResponse, name)
	}

	data["nameList"] = nameResponse
	data["wordsList"] = wordsResult
	appG.Response(http.StatusOK, e.SUCCESS, data)
	return
}
