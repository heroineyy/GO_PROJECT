package apis

import (
	"ginweb/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetUser(c *gin.Context) {
	var userName = c.Query("username")
	var password = c.Query("password")

	validate := validator.New()
	user1 := models.Users{UserName: userName, Password: password}
	err := validate.Struct(user1)
	if err != nil {
		c.JSON(200, gin.H{
			"code": false,
			"msg":  err.Error(),
		})
		return
	}
	success, errs := models.GetUser(userName, password)
	c.JSON(200, gin.H{
		"code": success,
		"msg":  errs,
	})
	//response.addHeader("access-control-allow-origin", "*");
}

func RegisterUser(c *gin.Context) {
	var userName = c.Request.FormValue("username")
	var password = c.Request.FormValue("password")
	validate := validator.New()
	user1 := models.Users{UserName: userName, Password: password}
	err := validate.Struct(user1)
	if err != nil {
		c.JSON(200, gin.H{
			"code": false,
			"msg":  err.Error(),
		})
		return
	}
	success, errs := models.RegisterUser(userName, password)
	c.JSON(200, gin.H{
		"code": success,
		"msg":  errs.Error(),
	})
}
