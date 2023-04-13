package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-name/src/pkg/app"
	"go-name/src/pkg/e"
	"net/http"
)

const (
	code2sessionURL = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	appID           = "wx7426614630ba44da"
	appSecret       = "你的AppSecret"
)

func sendWxAuthAPI(code string) (string, error) {
	url := fmt.Sprintf(code2sessionURL, appID, appSecret, code)
	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		return "", err
	}
	var wxMap map[string]string
	err = json.NewDecoder(resp.Body).Decode(&wxMap)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	return wxMap["openid"], nil
}

// GetOpenId 换取微信openid
func GetOpenId(c *gin.Context) {
	appG := app.Gin{C: c}

	code := c.Query("code")

	if code == "" {
		fmt.Print("get wx code error")
		appG.Response(http.StatusBadRequest, e.InvalidParams, "code is required")
		return
	}

	openid, error := sendWxAuthAPI(code)

	if error != nil {
		fmt.Print("sendWxAuthAPI code error", error)
		appG.Response(http.StatusBadRequest, e.ErrorGetWxOpenFail, "")
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, openid)
	return
}
