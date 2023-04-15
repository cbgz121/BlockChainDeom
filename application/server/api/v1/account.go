package v1

import (
	bc "application/blockchain"
	"application/pkg/app"
	"crypto/sha256"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var accountIds = map[string]bool{}

func QueryAccountList(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(AccountRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
        fmt.Println("9999999999999999",body.Username)
        fmt.Println("777777777777777777", body.Password)
	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.Username))
	bodyBytes = append(bodyBytes, []byte(body.Password))

	//调用智能合约
	resp, err := bc.ChannelQuery("queryAccountList", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	// 反序列化json
	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	fmt.Println(data)
//	if accountId, ok := data[0]["accountId"].(string); ok{
//		if accountIds[accountId] && len(data) == 1 {
//				appG.Response(http.StatusInternalServerError, "失败", "请注册")
//				return
//		}
//	}
	appG.Response(http.StatusOK, "成功", data)
}

func DeleteID(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(AccountRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	combined := body.Username + body.Password

	// 使用SHA-256算法进行hash
	hashed := sha256.Sum256([]byte(combined))
	// 将hash结果转换为16进制字符串
    hashStr := fmt.Sprintf("%x", hashed)
	hashStr = hashStr[:16]
	accountIds[hashStr] = true
	appG.Response(http.StatusOK, "成功", accountIds)
}
