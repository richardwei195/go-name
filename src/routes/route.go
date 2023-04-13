package routes

import (
	"github.com/gin-gonic/gin"
	"go-name/src/controllers"
)

func Routes(r *gin.Engine) {
	r.Use(gin.Logger())
	r.POST("/getNames", controllers.GetNames)
	r.GET("/prePay", controllers.GetPrePay)
	r.GET("/wxOpenId", controllers.GetOpenId)
}
