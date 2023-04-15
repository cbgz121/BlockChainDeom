package v1

import (
	"application/pkg/app"
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

type RegisterData struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

func Register(c *gin.Context) {
	// 创建数据库连接
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
        fmt.Println("11111111111111111111111")
	// 创建用户表格
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS uuu (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(50), password VARCHAR(50), email VARCHAR(30), phone VARCHAR(20), user_id VARCHAR(30))")
	if err != nil {
		panic(err.Error())
	}
	appG := app.Gin{C: c}
	reqBody := new(RegisterData)
	if err := c.ShouldBind(reqBody); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	// 检查用户是否已经存在
	row := db.QueryRow("SELECT id FROM uuu WHERE name = ?", reqBody.Username)
	var id int
	err = row.Scan(&id)
	if err != sql.ErrNoRows {
                fmt.Println("222222222222222")
		appG.Response(http.StatusBadRequest, "该用户名已经被注册", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
        fmt.Println("33333333333333")

	// 在数据库中插入用户
	_, err = db.Exec("INSERT INTO uuu (name, password, email, phone) VALUES (?, ?, ?, ?)", reqBody.Username, reqBody.Password, reqBody.Email, reqBody.Phone)
	if err != nil {
                fmt.Println("4444444444444444")
		appG.Response(http.StatusInternalServerError, "插入用户失败", fmt.Sprintf("参数出错%s", err.Error()))
                fmt.Println(err.Error())
		return
	}

	// 注册成功
	appG.Response(http.StatusOK, "注册成功", "注册成功")
}

