package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-name/src/pkg/app"
	"go-name/src/pkg/e"
	"go-name/src/pkg/pay"
	"net/http"
	"strconv"
)

func GetPrePay(c *gin.Context) {
	appG := app.Gin{C: c}

	amount, err := strconv.ParseInt(c.Query("amount"), 10, 64)

	if err != nil {
		fmt.Print("GetPrePay error", err)
		appG.Response(http.StatusBadRequest, e.InvalidParams, "amount is invalid")
		return
	}

	fmt.Print("GetPrePay query amount", amount)

	preResp, err := pay.PrePay(amount)

	if err != nil {
		appG.Response(http.StatusBadRequest, e.InvalidParams, "amount is invalid")
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, preResp)
	return
}
