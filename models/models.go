package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users struct {
	Name     string    `json:"name" bson:"name"`
	Email    string    `json:"email" bson:"email"`
	Mobile   string    `json:"mobile" bson:"mobile"`
	Password string    `json:"password" bson:"password"`
	Role     string    `json:"role" bson:"role"`
	Created  time.Time `json:"createdAt" bson:"createdAt"`
}

type SignInBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Mobile   string `json:"mobile"`
	Role     string `json:"role"`
}

type Session struct {
	ID      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email   string             `json:"email" bson:"email"`
	Expiry  time.Time          `json:"expiryAt" bson:"expiryAt"`
	Created time.Time          `json:"createdAt" bson:"createdAt"`
}
