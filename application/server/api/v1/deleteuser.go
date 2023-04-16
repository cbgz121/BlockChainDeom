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
type DeleteData struct {
	Id int `json:"id"`
}

func DeleteUser(c *gin.Context) {
	appG := app.Gin{C: c}
	// 创建数据库连接
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	reqBody := new(DeleteData)
	if err := c.ShouldBind(reqBody); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
        fmt.Println(reqBody.Id)
	// 删除指定ID的用户
	result, err := db.Exec("DELETE FROM uuu WHERE id = ?", reqBody.Id)
	if err != nil {
                fmt.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// 检查是否有任何行受到影响
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	} else if rowsAffected == 0 {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	// 返回删除成功的消息
	appG.Response(http.StatusOK, "删除成功", reqBody.Id)
}

