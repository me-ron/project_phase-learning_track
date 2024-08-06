package data

import (
	"context"
	"os"
	"task_manager/models"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = os.Getenv("JWT_SECRET")

func (tm *Taskmanager) Signup(user models.User) error{
	coll := tm.collection("users")
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	_, er := coll.InsertOne(context.TODO(), &user)
	if er != nil {
		return er
	}

	return nil
}


func (tm *Taskmanager) Login(user models.User) (string, error){
	coll := tm.collection("users")
	var userDB models.User
	err := coll.FindOne(context.TODO(), user).Decode(&userDB)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userDB.ID,
		"email":   userDB.Email,
		"isadmin":    userDB.IsAdmin,
	})
	jwtToken, err := token.SignedString(secretKey)
		if err != nil {
			return "", err
		}

	return jwtToken, nil
} 
