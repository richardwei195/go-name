package routes

import (
	"github.com/gin-gonic/gin"
	"go-name/src/controllers"
)

func Routes(r *gin.Engine) {
	r.POST("/getNames", controllers.GetNames)
}
