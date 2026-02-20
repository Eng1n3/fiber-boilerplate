package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	gs "github.com/randree/gormseeder"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	godotenv.Load()
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	fmt.Println("Connecting to database...", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{LogLevel: logger.Silent},
	)})
	if err != nil {
		fmt.Println(err.Error())
	}

	gs.InitSeeder(db, "seeders")
}
