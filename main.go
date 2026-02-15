package main

import (
	"fiber-boilerplate/app"
	"fiber-boilerplate/database"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	database.Connect()

	app := app.App()
	log.Fatal(app.Listen(":8080", fiber.ListenConfig{EnablePrefork: true}))
}
