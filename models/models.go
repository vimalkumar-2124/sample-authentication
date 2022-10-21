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
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email       string             `json:"email" bson:"email"`
	Expiry      int64              `json:"expiryAt" bson:"expiryAt"`
	Created     time.Time          `json:"createdAt" bson:"createdAt"`
	TokenString string             `json:"token" bson:"token"`
}

type ChangeUserPassword struct {
	Email        string `json:"email"`
	Old_Password string `json:"old_pass"`
	New_Password string `json:"new_pass"`
}

type TokenMetaData struct {
	Role   string `json:"role"`
	Expiry int64  `json:"expiryAt"`
}

type AllUser struct {
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Mobile  string    `json:"mobile"`
	Role    string    `json:"role"`
	Created time.Time `json:"createdAt"`
}
