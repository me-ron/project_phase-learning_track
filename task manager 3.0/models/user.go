package models

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserInput struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name,omitempty" bson:"name,omitempty"`
	Email     string             `json:"email,omitempty" bson:"email,omitempty"`
	Password  string			 `json:"password" bson:"password"`
	IsAdmin   bool 				 `json:"isadmin" bson:"isadmin"`
}


type DBUser struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name,omitempty" bson:"name,omitempty"`
	Email     string             `json:"email,omitempty" bson:"email,omitempty"`
	IsAdmin   bool 				 `json:"isadmin" bson:"isadmin"`
}

type Dclaims struct{
	jwt.StandardClaims
	ID primitive.ObjectID    
	Name string 
	Email string                  
	IsAdmin bool 
}

func ChangeToOutput(user UserInput) DBUser{
	var Ouser DBUser
	Ouser.ID = user.ID
	Ouser.Name = user.Name
	Ouser.Email = user.Email
	Ouser.IsAdmin = user.IsAdmin

	return Ouser
}