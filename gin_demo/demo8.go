package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//路由分组
func main() {
	router := gin.Default()

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.POST("/login", func(c *gin.Context) {
			fmt.Println("login")

		})
		v1.POST("/submit", func(c *gin.Context) {
			fmt.Println("submit")

		})
		v1.POST("/read", func(c *gin.Context) {
			fmt.Println("read")

		})
	}

	// Simple group: v2
	v2 := router.Group("/v2")
	{
		v2.POST("/login", func(c *gin.Context) {
			fmt.Println("login2")

		})
		v2.POST("/submit", func(c *gin.Context) {
			fmt.Println("submit2")

		})
		v2.POST("/read", func(c *gin.Context) {
			fmt.Println("read2")

		})
	}

	router.Run(":8080")
}
