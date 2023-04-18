package v1

import (
	"application/pkg/app"
	"database/sql"
	"net/http"
	"strconv"
	"strings"
        "fmt"

	"github.com/gin-gonic/gin"
)

// 定义一个结构体表示需要更新的数据
type UpData struct {
	ID       int64  `json:"id"`
	Name     string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Status   string `json:"status"`
}

// 定义一个函数，更新指定字段的值
func UpdateData(c *gin.Context) {
	appG := app.Gin{C: c}
	// 获取请求体中的数据
	var data UpData
	if err := c.ShouldBindJSON(&data); err != nil {
                fmt.Println(err)
		appG.Response(http.StatusBadRequest, "失败", err.Error())
		return
	}
        fmt.Println("111111111")

	// 连接数据库
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		appG.Response(http.StatusInternalServerError, "连接数据库失败", err.Error())
		return
	}
	defer db.Close()
        fmt.Println("111111111111")
	// 构建更新 SQL 语句
	var sqlBuilder strings.Builder
	sqlBuilder.WriteString("UPDATE uuu SET ")

	if data.Name != "" {
		sqlBuilder.WriteString("name = ?, ")
	}

	if data.Password != "" {
		sqlBuilder.WriteString("password = ?, ")
	}

	if data.Email != "" {
		sqlBuilder.WriteString("email = ?, ")
	}

	if data.Phone != "" {
		sqlBuilder.WriteString("phone = ?, ")
	}
	
  
	if data.Status != "" {
		sqlBuilder.WriteString("status = ?, ")
	}

	// 去除末尾的逗号和空格
	sqlStr := strings.TrimSuffix(sqlBuilder.String(), ", ")
        fmt.Println(sqlStr)
	// 添加 WHERE 子句
	sqlStr += " WHERE id = " + strconv.FormatInt(data.ID, 10)

	// 构建参数数组
	var params []interface{}
	if data.Name != "" {
		params = append(params, data.Name)
	}

	if data.Password != "" {
		params = append(params, data.Password)
	}

	if data.Email != "" {
		params = append(params, data.Email)
	}

	if data.Phone != "" {
		params = append(params, data.Phone)
	}

	if data.Status != "" {
		params = append(params, data.Status)
	}

	// 执行更新操作
	result, err := db.Exec(sqlStr, params...)
	if err != nil {
                fmt.Println(err)
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}

	// 获取更新的行数
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}

	// 返回更新成功的响应
	appG.Response(http.StatusOK, "更新成功", rowsAffected)
}

