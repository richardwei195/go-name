package routes

import (
	"github.com/gin-gonic/gin"
	"go-name/src/calendar"
	"log"
	"net/http"
	"strconv"
)

func GetDateResult(c *gin.Context) {
	year, err := strconv.Atoi(c.Query("year"))
	if err != nil {
		log.Printf("year type is not valid")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "year type is not valid",
		})
		return
	}

	month, err := strconv.Atoi(c.Query("month"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "month type is not valid",
		})
		return
	}

	day, err := strconv.Atoi(c.Query("day"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "day type is not valid",
		})
		return
	}
	hour, err := strconv.Atoi(c.Query("hour"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "hour type is not valid",
		})
		return
	}

	birthTime := calendar.DateReturn{
		Year:  year,
		Month: month,
		Day:   day,
		Hour:  hour,
	}.GetBirthTime()

	solarDate := calendar.GetSolarDate(birthTime)
	c.JSON(http.StatusOK, gin.H{
		"birthDate": birthTime,
		"solarDate": solarDate,
	})
}
