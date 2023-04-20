package v1

import (
	bc "application/blockchain"
	"application/pkg/app"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
        "database/sql"
        "crypto/sha256"

	"github.com/gin-gonic/gin"
)

type AuthenData struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name    string `json:"name"`
	IdCard  string  `json:"idcard"`
	Phone    string `json:"phone"`
}

func Authentication(c *gin.Context) {
	appG := app.Gin{C: c}
	reqBody := new(AuthenData)
	if err := c.ShouldBind(reqBody); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(reqBody.Username))
	bodyBytes = append(bodyBytes, []byte(reqBody.Password))
	bodyBytes = append(bodyBytes, []byte(reqBody.Name))
	bodyBytes = append(bodyBytes, []byte(reqBody.IdCard))
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
        updateUserByUsername(reqBody.Username,reqBody.Password)

	if len(data) != 0 {
		appG.Response(http.StatusOK, "成功", data)
	}
}

func updateUserByUsername(username string, password string) error {
	// 1. 连接数据库
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/block")
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}
	defer db.Close()
        
	// 2. 查询记录
	row := db.QueryRow("SELECT * FROM users WHERE username=?", username)
	var id int
	var pwd string
        var email string
        var phone string
        var status string
        var userid sql.NullString
	err = row.Scan(&id, &username, &pwd, &email, &phone, &status, &userid)
	if err != nil {
		return fmt.Errorf("failed to get user by username: %v", err)
	}
        
        combined := username + password

	// 使用SHA-256算法进行hash
	hashed := sha256.Sum256([]byte(combined))
	// 将hash结果转换为16进制字符串
        hashStr := fmt.Sprintf("%x", hashed)
	hashStr = hashStr[:16]

	// 3. 更新记录
	result, err := db.Exec("UPDATE users SET user_id=? WHERE id=?", hashStr, id)
	if err != nil {
		return fmt.Errorf("failed to update user ID: %v", err)
	}

	// 4. 检查更新是否成功
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no rows affected")
	}

	return nil
}
