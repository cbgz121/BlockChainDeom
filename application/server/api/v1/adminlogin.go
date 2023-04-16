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
type AdminData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func AdminLogin(c *gin.Context) {
	// 创建数据库连接
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// 创建用户表格
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS admin (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(50), password VARCHAR(50), admin_id VARCHAR(30))")
	if err != nil {
		panic(err.Error())
	}
	appG := app.Gin{C: c}
	reqBody := new(AdminData)
	if err := c.ShouldBind(reqBody); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

        // 在数据库中插入用户
//        _, err = db.Exec("INSERT INTO admin (name, password, admin_id) VALUES (?, ?, ?)", reqBody.Username, reqBody.Password, "ac0e7d037817094e")
//        if err != nil {
//                appG.Response(http.StatusInternalServerError, "插入用户失败", fmt.Sprintf("参数出错%s", err.Error()))
//                fmt.Println(err.Error())
//                return
//        }

	// 在数据库中查找用户
	row := db.QueryRow("SELECT name, password, admin_id FROM admin WHERE name = ?", reqBody.Username)
	var storedUsername, storedPassword string
        var storedUserid sql.NullString
	err = row.Scan(&storedUsername, &storedPassword, &storedUserid)
	if err != nil {
                fmt.Println(err)
		if err == sql.ErrNoRows {
			appG.Response(http.StatusUnauthorized, "用户不存在", fmt.Sprintf("参数出错%s", err.Error()))
			return
		}
		appG.Response(http.StatusInternalServerError, "查询用户失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	// 检查密码是否正确
	if reqBody.Password != storedPassword {
		appG.Response(http.StatusUnauthorized, "密码不正确", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var data []map[string]interface{}
        adminid := storedUserid.String
        if adminid == "" {
            adminid = "-11111"
        }
	data = append(data, map[string]interface{}{"accountId": adminid, "passWord": reqBody.Password, "userName": reqBody.Username})

	// 登录成功
	appG.Response(http.StatusOK, fmt.Sprintf("欢迎回来，%s！", storedUsername), data)
}

