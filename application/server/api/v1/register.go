package v1

import (
	"application/pkg/app"
	"fmt"
	"net/http"
	"bytes"
	"encoding/json"
	bc "application/blockchain"

	"github.com/gin-gonic/gin"
)

type RegisterData struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

func Register(c *gin.Context) {
	appG := app.Gin{C: c}
	reqBody := new(RegisterData)
	if err := c.ShouldBind(reqBody); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(reqBody.Username))
	bodyBytes = append(bodyBytes, []byte(reqBody.Password))
	bodyBytes = append(bodyBytes, []byte(reqBody.Email))
	bodyBytes = append(bodyBytes, []byte(reqBody.Phone))

	//调用智能合约
	resp, err := bc.ChannelExecute("register", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	var data map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	fmt.Println(data)

	// 假设处理完毕后，将数据存入数据库成功
	if len(data) != 0 {
		appG.Response(http.StatusOK, "成功", "注册成功")
	}
	// appG.Response(http.StatusOK, "成功", "注册成功")
}
