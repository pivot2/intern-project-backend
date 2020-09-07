package main

import (
	"package/products"
	"package/users"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20
	v1 := r.Group("/user")
	{
		v1.GET("/", users.Endpoint1)
		v1.GET("/:id", users.Endpoint2)
		v1.POST("/login", users.LoginEndpoint)
		v1.POST("/signup", users.SignUpEndpoint)
	}
	v2 := r.Group("/product")
	{
		v2.POST("/add", products.AddProductEndpoint)
		v2.GET("/:id", products.GetProductInfoRoute)
	}
	r.GET("/", func(c *gin.Context) {
		c.File("./test.html")
	})

	r.Run(":3000")
}
