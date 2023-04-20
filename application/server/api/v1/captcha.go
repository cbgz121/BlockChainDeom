package v1

import (
	"application/pkg/app"
	"net/http"
        "fmt"

	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
)

func GetCaptcha(c *gin.Context) {
	appG := app.Gin{C: c}
	captchaId := captcha.New()
//	appG.Response(http.StatusOK, gin.H{
//		"captchaId": captchaId,
//	})
        fmt.Println(88888888888888)

	captchaData := make(map[string]interface{})
	captchaData["captchaId"] = captchaId
	appG.Response(http.StatusOK, "成功", captchaData)
}
