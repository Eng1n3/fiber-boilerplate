package main

import (
	app "fiber-boilerplate"
	"fiber-boilerplate/database"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()
	database.Connect()
	app := app.App()
	log.Fatal(app.Listen(fmt.Sprintf("%s:%s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT"))))

}
