package main

import (
	"github.com/dgrijalva/jwt-go"
	"strings"
	"github.com/jinzhu/gorm"
	"os"
	"golang.org/x/crypto/bcrypt"
 //   "fmt"
)

/*
Структура прав доступа JWT
*/

func (account *Account) Validate() (map[string] interface{}, bool) {

	if strings.Count(account.Email, "@") != 1 {
		return Message(false, "Your email seems to be incorrect!"), false
	}

	if len(account.Password) < 6 {
		return Message(false, "Your password should contains more than 6 symbols!"), false
	}

	temp := &Account{}

	err := GetDB().Table("accounts").Where("email = ?", account.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return Message(false, "Connection error. Please retry"), false
	}
	if temp.Email != "" {
		return Message(false, "Email address already in use by another user."), false
	}
    
	return Message(false, "Requirement passed"), true
}

func (account *Account) Create() (map[string] interface{}) {

	if resp, ok := account.Validate(); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)
    
	GetDB().Create(account)

	if account.ID <= 0 {
		return Message(false, "Failed to create account, connection error.")
	}

	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString

	account.Password = ""

	response := Message(true, "Account has been created")
	response["account"] = account
	return response
}

func Login(email, password string) (map[string]interface{}) {

	account := &Account{}
	err := GetDB().Table("accounts").Where("email = ?", email).First(account).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return Message(false, "Email address not found")
		}
		return Message(false, "Connection error. Please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Пароль не совпадает!!
		return Message(false, "Invalid login credentials. Please try again")
	}

	account.Password = ""

	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString 
	
	resp := Message(true, "Logged In")
	resp["account"] = account
	return resp
}

func GetUser(u uint) *Account {

	acc := &Account{}
	GetDB().Table("accounts").Where("id = ?", u).First(acc)
	if acc.Email == "" {
		return nil
	}

	acc.Password = ""
	return acc
}
