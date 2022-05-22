package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userInfo struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func main() {
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		//username:=c.Query("username")
		//password :=c.Query("password")
		//u := userInfo{
		//	username:username,
		//	password:password,
		//}
		var u userInfo //声明一个结构体的变量
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}

	})
	r.POST("/post", func(c *gin.Context) {
		var u userInfo
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}

	})
	r.Run(":8081")
}
