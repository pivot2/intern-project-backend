package products

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Claims struct {
	UserID int
	Type   int
	jwt.StandardClaims
}

func AddProductEndpoint(c *gin.Context) {
	// var json Product
	// c.Bind(&json)

	// err := AddProduct(json)
	tknStr, _ := c.Cookie("token")
	fmt.Println(tknStr)
	var jwtKey = []byte("my_secret_key")
	claims := &Claims{}
	tkn, _ := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if !tkn.Valid {
		c.JSON(400, gin.H{
			"add": false,
			"err": "Invalid token",
		})
		return
	}
	var product Product
	product.SellerID = claims.UserID
	product.Brand = c.PostForm("brand")
	product.Categories = c.PostFormArray("categories")
	product.Name = c.PostForm("name")
	product.Price, _ = strconv.Atoi(c.PostForm("price")) //strconv.ParseFloat(c.PostForm("price"), 64)
	product.Size = c.PostFormArray("size")
	product.Colors = c.PostFormArray("colors")
	product.Quantity, _ = strconv.Atoi(c.PostForm("quantity"))
	product.Description = c.PostForm("des")
	form, _ := c.MultipartForm()
	files := form.File["file[]"]

	for i, file := range files {
		file.Filename = strconv.Itoa(i) + "_" + strconv.Itoa(int(time.Now().Unix())) + ".png"
		product.Images = append(product.Images, file.Filename)
		c.SaveUploadedFile(file, filepath.Join("../frontend/my-app/src/images", filepath.Base(file.Filename)))
	}
	err := AddProduct(product)
	if err != nil {
		c.JSON(400, gin.H{
			"add": false,
			"err": err,
		})
	} else {
		c.Redirect(http.StatusFound, "/products")
	}

}

func GetProductInfoRoute(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	c.JSON(200, gin.H{
		"info": GetProductInfo(id),
	})
}

func GetAllProductInfoRoute(c *gin.Context) {

	c.JSON(200, gin.H{
		"info": GetAllProductInfo(),
	})
}

func GetSellerProductInfoRoute(c *gin.Context) {
	fmt.Println("token")
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
		"info": GetAllProductofASeller(claims.UserID),
	})
}
