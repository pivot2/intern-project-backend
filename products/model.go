package products

import (
	"fmt"
	"package/db"
	"strconv"

	"github.com/lib/pq"
)

type Product struct {
	Id          int            `db:"product_id" json:"id"`
	SellerID    int            `db:"seller_id" json:"seller_id"`
	Name        string         `db:"name" json:"name"`
	Categories  pq.StringArray `db:"categories" json:"categories"`
	Brand       string         `db:"brand" json:"brand"`
	Price       int            `db:"price" json:"price"`
	Size        pq.StringArray `db:"size" json:"size"`
	Colors      pq.StringArray `db:"colors" json:"colors"`
	Quantity    int            `db:"quantity" json:"quantity"`
	Description string         `db:"description" json:"des"`
	Images      pq.StringArray `db:"image_link" json:"image"`
}

func AddProduct(product Product) error {
	sql := `Insert into product values (default,$1,$2,$3,
		$4,$5,$6,$7,$8,$9,$10)`
	_, err := db.Db.Query(sql, product.SellerID, product.Name,
		pq.Array(product.Categories), product.Brand, product.Price, pq.Array(product.Size),
		pq.Array(product.Colors), product.Quantity, product.Description, pq.Array(product.Images))

	if err != nil {
		panic(err)
	}
	return err
}

func GetProductInfo(id int) []Product {
	product := []Product{}
	fmt.Println(id)
	db.Db.Select(&product, "Select * from product where product_id="+strconv.Itoa(id))
	return product
}

func GetAllProductInfo() []Product {
	product := []Product{}
	db.Db.Select(&product, "Select * from product")
	return product
}

func GetAllProductofASeller(id int) []Product {
	product := []Product{}
	db.Db.Select(&product, "Select * from product where seller_id="+strconv.Itoa(id))
	fmt.Println(product)
	return product
}
