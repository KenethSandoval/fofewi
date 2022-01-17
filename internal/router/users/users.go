package users

import (
	_ "go.mongodb.org/mongo-driver/bson/primitive"
)

type Users struct {
	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
	Rol      string `bson:"rol" json:"rol"`
}
