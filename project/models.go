package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"fmt"
    "github.com/dgrijalva/jwt-go"
)

type Account struct {
	gorm.Model
	Email string `json:"email"`
	Password string `json:"password"`
	Token string `json:"token";sql:"-"`
}

type Attempt struct {
	gorm.Model
	Word string `json:"word"`
	Game string `json:"game"`
	Email string `json:"email"`
	Number string `json:"number"`
}

type AttemptAnswer struct {
	Word string `json:"word"`
	Game string `json:"game"`
	Email string `json:"email"`
	Number string `json:"number"`
}

type Game struct {
	gorm.Model
	Score int `json:"score"`
	Attempts int `json:"attempts"`
	Number string `json:"number"`
	Email string `json:"email"`
}

type GamesRequest struct {
	
	Email string `json:"email"`
}

type Token struct {
	UserId uint
	jwt.StandardClaims
}

var db *gorm.DB //база данных

func init() {

	e := godotenv.Load() //Загрузить файл .env
	if e != nil {
		fmt.Print(e)
	}

	username := "docker" //os.Getenv("db_user")
	password := "docker" //os.Getenv("db_pass")
	dbName := "docker"//os.Getenv("db_name")
	dbHost := "postgres"//os.Getenv("db_host")


	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=5432", dbHost, username, dbName, password) //Создать строку подключения
	fmt.Println(dbUri)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	db.Debug().AutoMigrate(&Account{}, &Game{}, &Attempt{}) //Миграция базы данных
}

// возвращает дескриптор объекта DB
func GetDB() *gorm.DB {
	return db
}
