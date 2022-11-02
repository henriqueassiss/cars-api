package main

import (
	"github.com/henriqueassiss/cars-api/database"
	"github.com/henriqueassiss/cars-api/routes"
	"github.com/joho/godotenv"
)

func main() {
	loadEnvVars()
	database.ConnectToDatabase()
	routes.Main()
}

func loadEnvVars() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
}
