package v1

import (
	bc "application/blockchain"
	"application/pkg/app"
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

func QueryAccountList(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(AccountRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
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
	// for i := range data {
	// 	if id, ok := data[i]["accountId"].(string); ok && len(id) >= 16 {
	// 		data[i]["accountId"] = id[:16]
	// 	}
	// }
	appG.Response(http.StatusOK, "成功", data)
}
