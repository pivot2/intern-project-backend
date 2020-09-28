package users

import (
	"strconv"

	"github.com/dgrijalva/jwt-go"
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
			"err":   "Your e-mail/password is invalid!",
			"login": false,
		})
		return
	}
	var jwtString string
	jwtString, sqlerr := LogIn(json.Username, json.Password)
	if sqlerr != "" {
		c.JSON(400, gin.H{
			"err":   sqlerr,
			"login": false,
		})
		return
	}
	tknStr := jwtString
	var jwtKey = []byte("my_secret_key")
	claims := &Claims{}
	jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	c.JSON(200, gin.H{
		"token":  jwtString,
		"type":   claims.Type,
		"userid": claims.UserID,
		"login":  true,
	})
	return

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
				"signup": false,
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
