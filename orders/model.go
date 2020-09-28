package orders

import (
	"fmt"
	"package/db"
	"strconv"
)

type Order struct {
	Id        int    `db:"id" json:"id"`
	Name      string `db:"name" json:"name"`
	Price     int    `db:"price" json:"price"`
	SellerID  int    `db:"seller_id" json:"seller_id"`
	ProductID int    `db:"product_id" json:"product_id"`
	Size      string `db:"size" json:"size"`
	Quantity  int    `db:"quantity" json:"quantity"`
	Status    string `db:"status" json:"status"`
}

func GetOrdersOfASeller(id int) []Order {
	fmt.Println("Id la", id)
	order := []Order{}
	db.Db.Select(&order, "select orders.*,product.name,product.price from orders left join product on orders.product_id=product.product_id where orders.seller_id="+strconv.Itoa(id))
	fmt.Println(order)
	return order
}

func SignUp(username string, password string, gmail string) string {
	var sqlerr string
	_, err := db.Db.Query("Insert into users values (default,$1,$2,$3,0)", username, password, gmail)
	if err != nil {
		sqlerr = "Account already exists"
	}
	return sqlerr
}
func AddOneOrder(order Order) error {
	fmt.Println(order.ProductID, order.SellerID, order.Quantity, order.Size, order.Status)
	_, err := db.Db.Query("Insert into orders values (default,$1,$2,$3,$4,$5)", order.ProductID, order.SellerID, order.Quantity, order.Size, order.Status)
	if err != nil {
		panic(err)
	}
	return err
}
