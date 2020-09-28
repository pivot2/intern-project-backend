package orders

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Claims struct {
	UserID int
	Type   int
	jwt.StandardClaims
}

func AddOrderEndpoint(c *gin.Context) {
	var order Order
	err := c.Bind(&order)
	if err != nil {
		c.JSON(400, gin.H{
			"err": "InvalidJSONBody",
		})
		return
	}
	error := AddOneOrder(order)
	if error != nil {
		c.JSON(400, gin.H{
			"error": error,
			"add":   false,
		})
	} else {
		c.JSON(200, gin.H{
			"add": true,
		})
	}

}
func GetSellerProductInfoRoute(c *gin.Context) {
	//fmt.Println("token")
	tknStr, _ := c.Cookie("token")
	fmt.Println(tknStr)
	var jwtKey = []byte("my_secret_key")
	claims := &Claims{}
	tkn, _ := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if !tkn.Valid {
		c.JSON(400, gin.H{
			"err": "Invalid token",
		})
		return
	}
	if claims.Type == 0 {
		c.JSON(400, gin.H{
			"err": "You a not a seller",
		})
		return
	}
	c.JSON(200, gin.H{
		"info": GetOrdersOfASeller(claims.UserID),
	})
}
