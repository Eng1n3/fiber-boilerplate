package app

import (
	"fiber-boilerplate/api/handlers"
	"fiber-boilerplate/api/routes"
	"fiber-boilerplate/database"
	"fiber-boilerplate/middleware"
	"fiber-boilerplate/pkg/auth"
	"fiber-boilerplate/pkg/config"
	"fiber-boilerplate/pkg/user"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/compress"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/helmet"
)

func App() *fiber.App {

	db := database.Connect()
	userRepo := user.NewRepository(db)

	app := fiber.New(config.FiberConfig())

	app.Use(middleware.LoggerConfig())
	app.Use(cors.New())
	app.Use(compress.New())
	app.Use(helmet.New())
	app.Use(middleware.TraceMiddleware)
	app.Use(middleware.RecoverConfig())

	// app.Route("/auth", routes.AuthRouter).Name("auth")

	jwtPrivateKey := os.Getenv("JWT_PRIVATE_KEY")
	jwtRefreshPrivateKey := os.Getenv("JWT_REFRESH_PRIVATE_KEY")
	jwtTTL, _ := strconv.Atoi(os.Getenv("JWT_TTL_IN_SECONDS"))
	jwtTTLInSeconds := time.Duration(jwtTTL) * time.Second
	jwtRefreshTTL, _ := strconv.Atoi(os.Getenv("JWT_REFRESH_TTL_IN_SECONDS"))
	jwtRefreshTTLInSeconds := time.Duration(jwtRefreshTTL) * time.Second
	authService := auth.NewService(userRepo, jwtPrivateKey, jwtRefreshPrivateKey, jwtTTLInSeconds, jwtRefreshTTLInSeconds)
	userService := user.NewService(userRepo)

	routes.AuthRouter(app, authService)
	routes.UserRouter(app, userService)

	// 404 Handler
	app.Use(handlers.NotFoundHandler())
	return app
}
