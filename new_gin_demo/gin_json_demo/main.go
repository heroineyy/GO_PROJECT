package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/json", func(c *gin.Context) {
		//方法1：使用map
		//data := map[string]interface{}{
		//	"name":    "小王子",
		//	"message": "hello world",
		//	"age":     18,
		//}
		data := gin.H{"message": "Hello world!"}
		c.JSON(http.StatusOK, data)
	})
	// 方法二：使用结构体,灵活使用tag来做订制化操作
	type msg struct {
		Name    string `json:"user"`
		Message string
		Age     int
	}

	r.GET("/new_json", func(c *gin.Context) {
		data := msg{
			"xyy",
			"hello",
			18,
		}
		c.JSON(http.StatusOK, data)
	})

	r.Run(":8081")
}
