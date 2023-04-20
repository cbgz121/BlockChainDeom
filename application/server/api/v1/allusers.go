package v1

import (
	"application/pkg/app"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
        Status   string    `json:"status"`
}

func AllUsers(c *gin.Context) {
// 连接数据库
		db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/block")
		if err != nil {
			panic(err)
		}
		defer db.Close()

		appG := app.Gin{C: c}
		// 查询所有用户
		rows, err := db.Query("SELECT id, username, password, email, phone, status FROM users")
		if err != nil {
			appG.Response(http.StatusInternalServerError, "查询失败", err.Error())
			return
		}
		defer rows.Close()

		// 构造用户列表
		users := []User{}
		for rows.Next() {
			user := User{}
			err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.Phone, &user.Status)
			if err != nil {
				appG.Response(http.StatusInternalServerError, "Scan失败", err.Error())
				return
			}
			users = append(users, user)
		}
		var data []map[string]interface{}
		for _, value := range users {
			data = append(data, map[string]interface{}{"ID": value.ID, "passWord": value.Password, "userName": value.Username, "email": value.Email, "phone": value.Phone, "status": value.Status})
		}

		// 返回用户列表
		appG.Response(http.StatusOK, "成功", data)

}
