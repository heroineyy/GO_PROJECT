package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/get", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "get",
		})
	})
	r.GET("/post", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "post",
		})
	})

	//any请求方法大杂烩
	r.Any("/test", func(c *gin.Context) {
		switch c.Request.Method {
		case "GET":
			c.JSON(http.StatusOK, gin.H{"method": "GET"})
		case http.MethodPost:
			c.JSON(http.StatusOK, gin.H{"method": "post"})
		}

	})

	//noroute
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"status": "not found"})
	})

	//路由组
	usergroup := r.Group("/user")
	{
		usergroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"method": "index"})
		})

		usergroup.GET("/login", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"method": "login"})
		})
	}

	r.Run(":8081")
}
