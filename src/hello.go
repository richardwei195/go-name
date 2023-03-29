package main

import (
	"fmt"
	"go-name/src/calendar"
	"net/http"

	"github.com/gin-gonic/gin"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello there!\n")
}
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": calendar.TestCalendar(),
		})
	})
	r.Run()
}
