// models/user.go
package models

import (
    "time"

    "go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
    ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
    Name      string            `json:"name" bson:"name"`
    Email     string            `json:"email" bson:"email"`
    Password  string            `json:"password" bson:"password"`
    CreatedAt time.Time         `json:"created_at" bson:"created_at"`
    UpdatedAt time.Time         `json:"updated_at" bson:"updated_at"`
}

type LoginInput struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}
