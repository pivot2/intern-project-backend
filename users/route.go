package users

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func Endpoint1(c *gin.Context) {
	c.JSON(200, gin.H{
		"users": GetAllUsers(),
	})
}

func Endpoint2(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	c.JSON(200, gin.H{
		"user": GetUserbyID(id),
	})
}

func LoginEndpoint(c *gin.Context) {
	var json Account
	err := c.Bind(&json)
	if err != nil {
		c.JSON(400, gin.H{
			"err":   "Thông tin đăng nhập không chính xác",
			"login": false,
		})
	} else {
		var user []Account
		user, sqlerr := LogIn(json.Username, json.Password)
		if sqlerr != "" {
			c.JSON(400, gin.H{
				"err":   sqlerr,
				"login": false,
			})
		} else {
			c.JSON(200, gin.H{
				"info":  user,
				"login": true,
			})
		}

	}
	// body := c.Request.Body
	// fmt.Println(body.username)

}

func SignUpEndpoint(c *gin.Context) {
	var json Account
	err := c.Bind(&json)
	if err != nil {
		c.JSON(400, gin.H{
			"error":  err,
			"signup": false,
		})
	} else {
		err := SignUp(json.Username, json.Password, json.Email)
		if err != "" {
			c.JSON(400, gin.H{
				"error":  err,
				"signup": true,
			})
		} else {
			c.JSON(200, gin.H{
				"signup": true,
			})
		}
	}
	// body := c.Request.Body
	// fmt.Println(body.username)

}
