package v1

import (
	"fmt"
	"application/pkg/app"
        "encoding/base64"
	"io/ioutil"
	"net/http"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func UpLoadImg(c *gin.Context) {
	// 初始化数据库连接
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/block")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// 创建数据库表
	err = createTable(db)
	if err != nil {
		panic(err.Error())
	}
	// 获取上传的文件
	file, err := c.FormFile("image")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("获取文件失败: %s", err.Error()))
		return
	}

	// 打开上传的文件
	src, err := file.Open()
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("打开文件失败: %s", err.Error()))
		return
	}
	defer src.Close()

	// 读取文件内容到字节数组中
	fileBytes, err := ioutil.ReadAll(src)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("读取文件内容失败: %s", err.Error()))
		return
	}

	// 存储文件到数据库
	err = storeFile(db,"test",fileBytes)
	// // 将字节数组存储到本地数据库中
	// err = ioutil.WriteFile(file.Filename, fileBytes, 0644)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("保存文件失败: %s", err.Error()))
		return
	}

	c.String(http.StatusOK, "文件上传成功")
}

// 存储文件到数据库
func storeFile(db *sql.DB, fileName string, fileBytes []byte) error {
	insertSQL := `
        INSERT INTO files (name, content) VALUES (?, ?)
    `
	_, err := db.Exec(insertSQL, fileName, fileBytes)
	if err != nil {
		return err
	}
	return nil
}

// 创建数据库表
func createTable(db *sql.DB) error {
	createTableSQL := `
        CREATE TABLE IF NOT EXISTS files (
            id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
            name VARCHAR(255) NOT NULL,
            content LONGBLOB NOT NULL
        )
    `
	_, err := db.Exec(createTableSQL)
	if err != nil {
		return err
	}
	return nil
}

func GetImg(c *gin.Context) {
	appG := app.Gin{C: c}
         // 初始化数据库连接
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

        var content []byte
	err = db.QueryRow("SELECT content FROM files WHERE name=?", "test").Scan(&content)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}

		// 将字节数组转换为 Base64 编码的字符串
	encoded := base64.StdEncoding.EncodeToString(content)
       // c.String(http.StatusOK, encoded)
       appG.Response(http.StatusOK, "cehnggong", encoded)
}
