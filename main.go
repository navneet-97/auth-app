// main.go
package main

import (
    "auth-app/config"
    "auth-app/handlers"
    "log"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "github.com/joho/godotenv"
)

func main() {
    // Load .env file
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

    // Connect to database
    config.ConnectDB()

    app := fiber.New()

    // Middleware
    app.Use(cors.New())
    app.Use(cors.New(cors.Config{
        AllowOrigins: "*",
        AllowHeaders: "Origin, Content-Type, Accept",
    }))

    // Routes
    app.Post("/register", handlers.Register)
    app.Post("/login", handlers.Login)

    log.Fatal(app.Listen(":3000"))
}