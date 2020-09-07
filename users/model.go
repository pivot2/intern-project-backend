package users

import (
	"database/sql"
	"fmt"
	"strconv"

	"package/db"
)

type Account struct {
	ID       int    `db:"user_id"`
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"password"`
	Email    string `db:"email" json:"gmail"`
	Type     int    `db:"type"`
}

func GetAllUsers() interface{} {
	Users := []Account{}
	db.Db.Select(&Users, "Select * from users")
	return &Users
}

func GetUserbyID(id int) []Account {
	User := []Account{}
	db.Db.Select(&User, "Select * from users where user_id="+strconv.Itoa(id))
	fmt.Println("Select * from users where user_id=?" + strconv.Itoa(id))
	return User
}

func LogIn(username string, password string) ([]Account, string) {
	User := []Account{}
	var sqlerr string
	err := db.Db.Select(&User, "Select * from users where username='"+username+"'")
	fmt.Println(err, " ", sql.ErrNoRows)
	if err != nil {
		panic(err)
	} else {

		if len(User) > 0 {
			if User[0].Password == password {

			} else {
				//err.Error() = "Wrong password"
				sqlerr = "Sai mật khẩu"
			}
		} else {
			sqlerr = "Tài khoản không tồn tại"
		}
	}

	//fmt.Println("Select * from users where user_id=?" + strconv.Itoa(id))
	return User, sqlerr
}

func SignUp(username string, password string, gmail string) string {
	var sqlerr string
	_, err := db.Db.Query("Insert into users values (default,$1,$2,$3,0)", username, password, gmail)
	if err != nil {
		sqlerr = "Tên tài khoản bị trùng"
	}
	return sqlerr
}
