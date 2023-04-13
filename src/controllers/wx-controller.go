package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	code2sessionURL = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	appID           = "你的AppID"
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

func GetOpenId(c *gin.Context) {

}
