package database

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectToDatabase() {
	h := os.Getenv("DATABASE_HOST")
	u := os.Getenv("DATABASE_USER")
	pa := os.Getenv("DATABASE_PASSWORD")
	n := os.Getenv("DATABASE_NAME")
	po := os.Getenv("DATABASE_PORT")

	if h == "" || u == "" || pa == "" || n == "" || po == "" {
		panic("You must set your database environmental variables corretly")
	}

	dsn := "host=" + h + " user=" + u + " password=" + pa + " dbname=" + n + " port=" + po + " sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
}
