package users

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"teamsy/internal/pkg/db"
	"time"
)

type User struct {
	gorm.Model

	Username string `gorm: unique`
	Email    string `gorm: unique`
	Password string
	Birthday time.Time
}

func (user User) Save() (uint, error) {

	hashPass, err := HashPassword(user.Password)
	if err != nil {
		return 0, err
	}

	userDB := User{
		Model:    gorm.Model{},
		Username: user.Username,
		Email:    user.Email,
		Password: hashPass,
		Birthday: user.Birthday,
	}

	res := db.Db.Create(&userDB)

	return userDB.ID, res.Error
}

func GetUserIdByUsername(username string) (int, error) {

	var user User

	db.Db.Where("username = ?", username).First(&user)

	if user.Username == "" {
		return 0, errors.New("there is no user by this username")
	}

	return int(user.ID), nil

}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func ParseBirthdayDate(date string) (time.Time, error) {
	birth, err := time.Parse("2006-01-02", date)
	if err != nil {
		log.Println(err)
		return birth, err
	}

	return birth, nil
}
