package v1

import (
	"application/pkg/app"
	"database/sql"
	"net/http"
	"strings"
        "fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// 定义一个结构体来表示搜索请求
type SearchRequest struct {
	Keyword string `json:"keyword"`
}

type Userkey struct {
	ID       int
	Name     string
	PassWord string
	Email    string
	Phone    string
        Status   string
}

// 定义搜索函数
func KeyWordSearch(c *gin.Context) {
	appG := app.Gin{C: c}
	// 获取请求中的关键字
	var req SearchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		appG.Response(http.StatusBadRequest, "json失败", err.Error())
		return
	}

	// 打开数据库连接
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/block")
	if err != nil {
		appG.Response(http.StatusInternalServerError, "连接数据库失败", err.Error())
		return
	}
	defer db.Close()

	// 获取表中的所有列
	columns, err := getColumns(db)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "获取表中列失败", err.Error())
		return
	}
	columns = columns[:len(columns)-2]

	// 构造查询语句
	var conditions []string
	for _, col := range columns {
		conditions = append(conditions, col+" LIKE ?")
	}
	query := "SELECT * FROM users WHERE " + strings.Join(conditions, " OR ")

	// 搜索数据
	var args []interface{}
	for range columns {
		args = append(args, "%"+req.Keyword+"%")
	}
	rows, err := db.Query(query, args...)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "查询失败", err.Error())
		return
	}
	defer rows.Close()

	// 读取查询结果到结构体数组中
	var users []Userkey 
	for rows.Next() {
		var user Userkey 
		var storedUserid sql.NullString
		err := rows.Scan(&user.ID, &user.Name, &user.PassWord, &user.Email, &user.Phone, &user.Status, &storedUserid)
		if err != nil {
			appG.Response(http.StatusInternalServerError, "读取查询结果失败", err.Error())
			return
		}
		users = append(users, user)
	}

	// 检查是否有错误发生
	if err := rows.Err(); err != nil {
		appG.Response(http.StatusInternalServerError, "查询结果错误", err.Error())
		return
	}

	var data []map[string]interface{}
	for _, v := range users {
		data = append(data, map[string]interface{}{"ID": v.ID, "userName": v.Name, "passWord": v.PassWord, "email": v.Email, "phone": v.Phone, "status": v.Status})
	}
        fmt.Println(data)

	// 将结果返回给客户端
	appG.Response(http.StatusOK, "查询成功", data)

	// 返回结果
	// appG.Response(http.StatusOK, "成功", result)
}

// 获取表中的所有列
func getColumns(db *sql.DB) ([]string, error) {
	var columns []string

	rows, err := db.Query("SELECT COLUMN_NAME FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ?", "block", "users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var column string
		if err := rows.Scan(&column); err != nil {
			return nil, err
		}
		columns = append(columns, column)
	}

	return columns, nil
}

