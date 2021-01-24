package db

import (
	"fmt"
	"gin_sample/entity"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"os"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	err = godotenv.Load(fmt.Sprintf(".env.%s", os.Getenv("GO_ENV")))
	if err != nil {
		panic(err)
	}
	db, err = gorm.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_NAME"),
		),
	)
	if err != nil {
		log.Fatalln(err)
	}
	autoMigration()
}

func autoMigration() {
	db.AutoMigrate(&entity.User{})
}

func GetDb() *gorm.DB {
	return db
}

func Close() {
	if err = db.Close(); err != nil {
		log.Fatalln(err)
	}
}
