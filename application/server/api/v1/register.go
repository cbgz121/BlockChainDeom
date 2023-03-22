package v1

import (
	"application/pkg/app"
	"crypto/sha256"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"fabric-realty/chaincode/pkg/utils"
	"fabric-realty/chaincode/model"
	"github.com/hyperledger/fabric/core/chaincode/shim"
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
	if err := c.BindJSON(&reqBody); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	// 在这里可以对前端传来的数据进行处理，例如将密码加密后存入数据库等等
	// 将用户名和密码合并成一个字符串
	combined := reqBody.Username + reqBody.Password

	// 使用SHA-256算法进行hash
	hashed := sha256.Sum256([]byte(combined))
	// 将hash结果转换为16进制字符串
    hashStr := fmt.Sprintf("%x", hashed)

	account := &model.Account{
		AccountId: hashStr,
		UserName:  reqBody.Username,
		Balance:   5000000,
	}
	
	// 写入账本
	if err := utils.WriteLedger(account, stub, model.AccountKey, []string{hashStr}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}

	// 假设处理完毕后，将数据存入数据库成功
	appG.Response(http.StatusOK, "成功", "注册成功")
}
