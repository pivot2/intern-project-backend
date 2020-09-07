package products

import (
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func AddProductEndpoint(c *gin.Context) {
	// var json Product
	// c.Bind(&json)

	// err := AddProduct(json)
	var product Product
	product.SellerID, _ = strconv.Atoi(c.PostForm("seller_id"))
	product.Brand = c.PostForm("brand")
	product.Categories = c.PostFormArray("categories")
	product.Name = c.PostForm("name")
	product.Price, _ = strconv.Atoi(c.PostForm("price"))
	product.Size = c.PostFormArray("size")
	product.Colors = c.PostFormArray("colors")
	product.Quantity, _ = strconv.Atoi(c.PostForm("quantity"))
	product.Description = c.PostForm("des")
	form, _ := c.MultipartForm()
	files := form.File["file[]"]

	//fmt.Println(filepath.Join("upload", "test.png"))
	for i, file := range files {
		file.Filename = strconv.Itoa(i) + "_" + strconv.Itoa(int(time.Now().Unix())) + ".png"
		product.Images = append(product.Images, file.Filename)
		c.SaveUploadedFile(file, filepath.Join("upload", filepath.Base(file.Filename)))
	}

	c.JSON(200, gin.H{
		"add":  true,
		"info": product,
	})

}

func GetProductInfoRoute(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	c.JSON(200, gin.H{
		"info": GetProductInfo(id),
	})
}
