# Auth App

This is a simple web application built using the Go Fiber framework. The application allows users to register and log in.

## Features

- **User Registration**: Users can register with their name, email, and password.
- **Password Hashing**: Passwords are securely hashed before being stored in the database.
- **User Login**: Users can log in using their email and password.
- **JWT Token**: A success message or JWT token is returned upon successful login.
- **Error Handling**: The application validates input fields and handles errors such as incorrect credentials and missing fields.

## Prerequisites

Make sure you have the following installed:

- Go (version 1.18 or later)
- MongoDB Atlas account (for database)
- A `.env` file for storing sensitive information like database URI and JWT secret.

## Environment Variables

Create a `.env` file in the root of the project directory and add the following environment variables:
  // .env
  MONGODB_URI=mongodb://localhost:27017
  DB_NAME=auth_db
  JWT_SECRET=your-secret-key

## Installation & Setup

1. **Create Project Directory**:

    ```bash
    mkdir auth-app
    cd auth-app
    ```

2. **Initialize Go Module**:

    ```bash
    go mod init auth-app
    ```

3. **Install Dependencies**:

    Run the following command to install the necessary dependencies:

    ```bash
    go get github.com/gofiber/fiber/v2
    go get github.com/joho/godotenv
    go get go.mongodb.org/mongo-driver/mongo
    go get golang.org/x/crypto/bcrypt
    go get github.com/golang-jwt/jwt/v5
    ```

## Running the Application

To run the application, use the following command:

```bash
go run main.go
