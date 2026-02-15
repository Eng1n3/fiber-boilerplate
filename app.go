package main

import (
	"fiber-boilerplate/database"
	"fiber-boilerplate/routes"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/compress"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/helmet"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	database.Connect()

	app := fiber.New()

	app.Use(cors.New())
	app.Use(compress.New())
	app.Use(helmet.New())

	app.Route("/auth", routes.AuthRoute).Name("auth")
	app.Route("/user", routes.UserRoute).Name("user")
	log.Fatal(app.Listen(fmt.Sprintf(":%s", os.Getenv("APP_PORT"))))
}
