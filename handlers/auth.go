// handlers/auth.go
package handlers

import (
    "auth-app/config"
    "auth-app/models"
    "auth-app/utils"
    "context"
    "os"
    "time"

    "github.com/gofiber/fiber/v2"
    "github.com/golang-jwt/jwt/v5"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

func Register(c *fiber.Ctx) error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    user := new(models.User)
    if err := c.BodyParser(user); err != nil {
        return c.Status(400).JSON(fiber.Map{
            "error": "Invalid input",
        })
    }

    // Validate input
    if user.Email == "" || user.Password == "" || user.Name == "" {
        return c.Status(400).JSON(fiber.Map{
            "error": "Missing required fields",
        })
    }

    // Check if email already exists
    collection := config.DB.Collection("users")
    var existingUser models.User
    err := collection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&existingUser)
    if err == nil {
        return c.Status(400).JSON(fiber.Map{
            "error": "Email already exists",
        })
    }

    // Hash password
    hashedPassword, err := utils.HashPassword(user.Password)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{
            "error": "Couldn't hash password",
        })
    }

    // Prepare user document
    now := time.Now()
    user.Password = hashedPassword
    user.CreatedAt = now
    user.UpdatedAt = now

    // Insert user
    result, err := collection.InsertOne(ctx, user)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{
            "error": "Couldn't create user",
        })
    }

    user.ID = result.InsertedID.(primitive.ObjectID)

    return c.Status(201).JSON(fiber.Map{
        "message": "User created successfully",
        "user_id": user.ID.Hex(),
    })
}

func Login(c *fiber.Ctx) error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    input := new(models.LoginInput)
    if err := c.BodyParser(input); err != nil {
        return c.Status(400).JSON(fiber.Map{
            "error": "Invalid input",
        })
    }

    // Find user
    collection := config.DB.Collection("users")
    var user models.User
    err := collection.FindOne(ctx, bson.M{"email": input.Email}).Decode(&user)
    if err != nil {
        return c.Status(400).JSON(fiber.Map{
            "error": "User not found",
        })
    }

    // Check password
    if !utils.CheckPasswordHash(input.Password, user.Password) {
        return c.Status(400).JSON(fiber.Map{
            "error": "Invalid password",
        })
    }

    // Generate JWT token
    token := jwt.New(jwt.SigningMethodHS256)
    claims := token.Claims.(jwt.MapClaims)
    claims["user_id"] = user.ID.Hex()
    claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

    t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
    if err != nil {
        return c.Status(500).JSON(fiber.Map{
            "error": "Couldn't generate token",
        })
    }

    return c.JSON(fiber.Map{
        "token": t,
    })
}