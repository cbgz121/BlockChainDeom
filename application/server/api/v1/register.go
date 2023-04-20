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
	Status   string `json:"status"`
}

func Register(c *gin.Context) {
	// 创建数据库连接
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/block")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("11111111111111111111111")
	// 创建用户表格
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INT AUTO_INCREMENT PRIMARY KEY, username VARCHAR(25), password VARCHAR(50), email VARCHAR(30), phone VARCHAR(20), status VARCHAR(1), user_id VARCHAR(30))")
	if err != nil {
		panic(err.Error())
	}
	appG := app.Gin{C: c}
	reqBody := new(RegisterData)
	if err := c.ShouldBind(reqBody); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "请求出错"})
		// appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	// 检查用户是否已经存在
	row := db.QueryRow("SELECT id FROM users WHERE username = ?", reqBody.Username)
	var id int
	err = row.Scan(&id)
	if err != sql.ErrNoRows {
		fmt.Println("222222222222222")
        c.JSON(http.StatusBadRequest, gin.H{"error": "该用户名已经被注册"})
		// appG.Response(http.StatusBadRequest, "该用户名已经被注册", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	fmt.Println("33333333333333")

	// 在数据库中插入用户
	_, err = db.Exec("INSERT INTO users (username, password, email, phone, status) VALUES (?, ?, ?, ?, ?)", reqBody.Username, reqBody.Password, reqBody.Email, reqBody.Phone, reqBody.Status)
	if err != nil {
		fmt.Println("4444444444444444")
        c.JSON(http.StatusBadRequest, gin.H{"error": "插入用户失败"})
		// appG.Response(http.StatusInternalServerError, "插入用户失败", fmt.Sprintf("参数出错%s", err.Error()))
		// fmt.Println(err.Error())
		return
	}

	// 注册成功
	appG.Response(http.StatusOK, "注册成功", "注册成功")
}

