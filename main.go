package main

import (
    "auth-app/config"
    "auth-app/handlers"
    "log"
    "path/filepath"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "github.com/joho/godotenv"
)

func main() {
    // Load .env file
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

    // Connect to the database
    config.ConnectDB()

    app := fiber.New()

    // Middleware
    app.Use(cors.New())
    app.Use(cors.New(cors.Config{
        AllowOrigins: "*",
        AllowHeaders: "Origin, Content-Type, Accept",
    }))

    // Serve static files (HTML, CSS, JS, etc.)
    app.Static("/", "./public") // Serve static files from the "public" folder

    // Routes to serve HTML files
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendFile(filepath.Join("./public", "index.html")) // Serve index.html
    })
    app.Get("/login", func(c *fiber.Ctx) error {
        return c.SendFile(filepath.Join("./public", "login.html")) // Serve login.html
    })
    app.Get("/register", func(c *fiber.Ctx) error {
        return c.SendFile(filepath.Join("./public", "register.html")) // Serve register.html
    })
    app.Get("/dashboard", func(c *fiber.Ctx) error {
        return c.SendFile(filepath.Join("./public", "dashboard.html")) // Serve dashboard.html
    })

    // API Routes for user registration and login
    app.Post("/register", handlers.Register)
    app.Post("/login", handlers.Login)

    log.Fatal(app.Listen(":3000"))
}
