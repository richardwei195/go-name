package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-name/src/calendar"
	"net/http"
	"strconv"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello there!\n")
}
func main() {
	r := gin.Default()
	r.GET("/dateResult", func(c *gin.Context) {
		year, err := strconv.Atoi(c.Query("year"))
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "year type is not valid",
			})
			return
		}

		month, err := strconv.Atoi(c.Query("month"))

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "month type is not valid",
			})
			return
		}

		day, err := strconv.Atoi(c.Query("day"))

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "day type is not valid",
			})
			return
		}
		hour, err := strconv.Atoi(c.Query("hour"))

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
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
	})
	r.Run()
}
