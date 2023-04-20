package v1

import (
       "application/pkg/app"
	"database/sql"
	"fmt"
	"net/http"
       _  "errors"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// 解析请求参数
type UserData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	// 创建数据库连接
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/block")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
        
        appG := app.Gin{C: c}
	reqBody := new(UserData)
	if err := c.ShouldBind(reqBody); err != nil {
                c.JSON(http.StatusBadRequest, gin.H{"error": "请求出错"})
		return
	}
	// 在数据库中查找用户
	row := db.QueryRow("SELECT username, password, status, user_id FROM users WHERE username = ?", reqBody.Username)
	var storedUsername, storedPassword, storedStatus string
	var storedUserid sql.NullString
	err = row.Scan(&storedUsername, &storedPassword, &storedStatus, &storedUserid)
	if err != nil {
		if err == sql.ErrNoRows {
                        c.JSON(http.StatusUnauthorized, gin.H{"error": "用户不存在"})
			return
		}
                c.JSON(http.StatusInternalServerError, gin.H{"error": "查询用户失败"})
		return
	}

	// 检查密码是否正确
	if reqBody.Password != storedPassword {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "密码错误"})
		return
	}
	var data []map[string]interface{}
	userid := storedUserid.String
	if userid == "" {
		userid = "-11111"
	}
	data = append(data, map[string]interface{}{"accountId": userid, "passWord": reqBody.Password, "userName": reqBody.Username})
	fmt.Println(data)
    if storedStatus == "0" {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "账户被禁用"})
        return
    }

	// 登录成功
	appG.Response(http.StatusOK, "成功", data)
}

