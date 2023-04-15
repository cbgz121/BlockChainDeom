package v1

import (
	"application/pkg/app"
	"database/sql"
	"fmt"
	"net/http"

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
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	appG := app.Gin{C: c}
	reqBody := new(UserData)
	if err := c.ShouldBind(reqBody); err != nil {
		fmt.Println("11111111111111111111111")
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	// 在数据库中查找用户
	row := db.QueryRow("SELECT name, password, user_id FROM uuu WHERE name = ?", reqBody.Username)
	var storedUsername, storedPassword string
        var storedUserid sql.NullString
        err = row.Scan(&storedUsername, &storedPassword, &storedUserid)
        if err != nil {
                if err == sql.ErrNoRows {
                        fmt.Println("2222222222222222222222222")
                        appG.Response(http.StatusUnauthorized, "用户不存在",fmt.Sprintf("参数出错%s", err.Error()))
                        return
                }
                fmt.Println(storedUsername)
                fmt.Println(storedPassword)
                fmt.Println(storedUserid)
                fmt.Println(err)
                appG.Response(http.StatusInternalServerError, "查询用户失败", fmt.Sprintf("参数出错%s", err.Error()))
                return
        }

	// 检查密码是否正确
        if reqBody.Password != storedPassword {
                fmt.Println("4444444444444444444444444")
                appG.Response(http.StatusUnauthorized, "密码不正确", fmt.Sprintf("参数出错%s", err.Error()))
                return
        }
        fmt.Println("555555555555555555555")
	var data []map[string]interface{}
        userid := storedUserid.String
        if userid == "" {
            userid = "-11111"
        }
	data = append(data, map[string]interface{}{"accountId": userid, "passWord": reqBody.Password, "userName": reqBody.Username})
	fmt.Println(data)
	

	// 登录成功
	appG.Response(http.StatusOK, fmt.Sprintf("欢迎回来，%s！", reqBody.Username), data)
}

