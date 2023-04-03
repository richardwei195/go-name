package controllers

import (
	"github.com/gin-gonic/gin"
	"go-name/src/calendar"
	"go-name/src/config"
	"go-name/src/models"
	"gorm.io/gorm"
	"net/http"
)

var (
	db *gorm.DB = config.ConnectDB()
)

type NameRequest struct {
	Year    int `json:"year"`
	Month   int `json:"month"`
	Day     int `json:"day"`
	Hour    int `json:"hour"`
	SurName int `json:"surName"`
}

// GetNames 获取用户名字分析
func GetNames(c *gin.Context) {
	var body NameRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	birthTime := calendar.DateReturn{
		Year:  body.Year,
		Month: body.Month,
		Day:   body.Day,
		Hour:  body.Hour,
	}.GetBirthTime()

	var nameList []models.TName

	error := db.Find(&nameList)

	if error.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "nameResult is null",
		})
		return
	}
	solarDate := calendar.GetSolarDate(birthTime)
	c.JSON(http.StatusOK, gin.H{
		"birthDate":  birthTime,
		"solarDate":  solarDate,
		"nameResult": nameList,
	})
}
