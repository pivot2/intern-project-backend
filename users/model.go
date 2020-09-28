package users

import (
	"package/db"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Account struct {
	ID       int    `db:"user_id"`
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"password"`
	Email    string `db:"email" json:"gmail"`
	Type     int    `db:"type"`
}

type Claims struct {
	UserID int
	Type   int
	jwt.StandardClaims
}

func GetAllUsers() []Account {
	Users := []Account{}
	db.Db.Select(&Users, "Select * from users")
	return Users
}

func GetUserbyID(id int) []Account {
	User := []Account{}
	db.Db.Select(&User, "Select * from users where user_id="+strconv.Itoa(id))
	return User
}

func LogIn(username string, password string) (string, string) {
	User := []Account{}
	var sqlerr string
	err := db.Db.Select(&User, "Select * from users where email='"+username+"'")
	if err != nil {
		panic(err)
	} else {
		if len(User) > 0 {
			if User[0].Password == password {
				var jwtKey = []byte("my_secret_key")
				expirationTime := time.Now().Add(120 * time.Minute)
				claims := &Claims{
					UserID: User[0].ID,
					Type:   User[0].Type,
					StandardClaims: jwt.StandardClaims{

						ExpiresAt: expirationTime.Unix(),
					},
				}
				token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
				jwtString, _ := token.SignedString(jwtKey)
				return jwtString, ""
			} else {
				//err.Error() = "Wrong password"
				sqlerr = "Wrong Password"
				return "", sqlerr
			}
		} else {
			sqlerr = "Wrong Gmail"
			return "", sqlerr
		}
	}

	//fmt.Println("Select * from users where user_id=?" + strconv.Itoa(id))

}

func SignUp(username string, password string, gmail string) string {
	var sqlerr string
	_, err := db.Db.Query("Insert into users values (default,$1,$2,$3,0)", username, password, gmail)
	if err != nil {
		sqlerr = "Email already been used"
	}
	return sqlerr
}
