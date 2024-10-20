package router

import (
	"fmt"
	"go-admin/entery"

	"github.com/gin-gonic/gin"
)

func InitRouterIndex() {
	// 注册路由
	r := gin.Default()
	r.Use(cors())

	// 更改为发布模式
	gin.SetMode(gin.ReleaseMode)

	// =========================================
	// 静态资源
	r.Static("/r", "./dist")

	// =========================================
	// 注册路由
	main := r.Group("/api/v50")
	{
		InitRouterGroupUser(main) // 注册用户模块路由
	}

	r.Run(fmt.Sprintf(":%s", entery.JsonEntery.ServicePort))
}

// 解决跨域
func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
