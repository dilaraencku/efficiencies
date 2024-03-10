package db

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
)

var MyDB *gorm.DB

func DB() *gorm.DB {
	return MyDB
}

type Database struct {
	Host     string `envgonfig:"DATABASE_HOST" required:"true"`
	Port     int    `envgonfig:"DATABASE_PORT" required:"true"`
	User     string `envgonfig:"DATABASE_USER" required:"true"`
	Password string `envgonfig:"DATABASE_PASSWORD" required:"true"`
	Name     string `envgonfig:"DATABASE_NAME" required:"true"`
}

var e error

func Init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Print("Error loading .env file : ", err)
	}

	host := os.Getenv("DATABASE_HOST")
	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")
	name := os.Getenv("DATABASE_NAME")
	portStr := os.Getenv("DATABASE_PORT")

	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Printf("Port numarasını çevirme hatası: %v", err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", host, user, password, name, port)

	MyDB, e = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if e != nil {
		panic(e)
	}
}
