package routers

import (
	v1 "application/api/v1"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}

// InitRouter 初始化路由信息
func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(Cors())

	apiV1 := r.Group("/api/v1")
	{
		apiV1.GET("/hello", v1.Hello)
                apiV1.GET("/captcha", v1.GetCaptcha)
                apiV1.POST("/login", v1.Login)
		apiV1.POST("/register",v1.Register)
		apiV1.POST("/deleteid",v1.DeleteID)
                apiV1.POST("/authentication", v1.Authentication)
		apiV1.POST("/queryAccountList", v1.QueryAccountList)
		apiV1.POST("/createRealEstate", v1.CreateRealEstate)
		apiV1.POST("/queryRealEstateList", v1.QueryRealEstateList)
		apiV1.POST("/createSelling", v1.CreateSelling)
		apiV1.POST("/createSellingByBuy", v1.CreateSellingByBuy)
		apiV1.POST("/querySellingList", v1.QuerySellingList)
		apiV1.POST("/querySellingListByBuyer", v1.QuerySellingListByBuyer)
		apiV1.POST("/updateSelling", v1.UpdateSelling)
		apiV1.POST("/createDonating", v1.CreateDonating)
		apiV1.POST("/queryDonatingList", v1.QueryDonatingList)
		apiV1.POST("/queryDonatingListByGrantee", v1.QueryDonatingListByGrantee)
		apiV1.POST("/updateDonating", v1.UpdateDonating)
	}
        r.Static("/captcha", "./captcha")
	return r
}
