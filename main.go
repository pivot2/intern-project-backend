package main

import (
	"package/orders"
	"package/products"
	"package/users"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20
	g1 := r.Group("/user")
	{
		g1.GET("/", users.Endpoint1)
		g1.GET("/:id", users.Endpoint2)
		g1.POST("/login", users.LoginEndpoint)
		g1.POST("/signup", users.SignUpEndpoint)
	}
	g2 := r.Group("/product")
	{
		g2.POST("/add", products.AddProductEndpoint)
		g2.GET("/:id/admin", products.GetSellerProductInfoRoute)
		//g2.GET("/show/:id", products.GetSellerProductInfoRoute)
		g2.GET("/:id", products.GetProductInfoRoute)
		g2.GET("/", products.GetAllProductInfoRoute)

	}
	g3 := r.Group("/order")
	{
		g3.GET("/:id", orders.GetSellerProductInfoRoute)
		g3.POST("/add", orders.AddOrderEndpoint)
	}
	r.Run(":8080")
}
