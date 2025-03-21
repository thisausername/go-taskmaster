package server

import (
	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由和中间件
func InitRouter() *gin.Engine {
	//使用Gin的默认引擎(已集成了Logger和Recovery中间件)
	router := gin.Default()
	//添加通用中间件
	router.Use(
		//CORS中间件示例
		func(c *gin.Context) {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Next()
		},
	)

	apiGroup := router.Group("/api/v1")
	{
		apiGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status": "ok",
			})
		})
	}

	return router
}

// listTask 示例控制器
func listTask(c *gin.Context) {
	//查询数据库
	c.JSON(200, gin.H{"tasks": []string{"task1", "task2"}})
}

func createTask(c *gin.Context) {
	c.JSON(201, gin.H{
		"id": "123",
	})
}
