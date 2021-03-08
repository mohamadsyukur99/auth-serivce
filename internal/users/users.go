package users

import (
	"log"

	database "github.com/mohamadsyukur99/auth-service/internal/pkg/db/mysql"
	"golang.org/x/crypto/bcrypt"
)

// User struct
type User struct {
	ID       string `json="id"`
	Username string `json="username"`
	Password string `json="password"`
}

// Create ...
func (user *User) Create() {

	sql := "INSERT INTO Users(Username,Password) VALUES(?,?)"
	statement, err := database.Db.Prepare(sql)
	print(statement)
	if err != nil {
		log.Fatal(err)
	}
	// memastikan statement ditutup setelah selesai
	defer statement.Close()

	hashedPassword, err := HashPassword(user.Password)
	_, err = statement.Exec(user.Username, hashedPassword)
	if err != nil {
		log.Fatal(err)
	}
}

//HashPassword hashes given password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
