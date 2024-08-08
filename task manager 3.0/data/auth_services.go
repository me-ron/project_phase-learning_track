package data

import (
	"context"
	"errors"
	"os"
	"task_manager/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = []byte(os.Getenv("JWT_SECRET"))

func (tm *Taskmanager) Signup(user models.UserInput) (models.UserInput, error) {
	coll := tm.collection("users")

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.UserInput{}, err
	}
	user.Password = string(hashedPassword)

	// Set other user properties
	user.ID = primitive.NewObjectID()
	user.IsAdmin = false

	// Insert the new user
	_, er := coll.InsertOne(context.TODO(), &user)
	if er != nil {
		// Check if the error is due to a duplicate key
		if mongo.IsDuplicateKeyError(er) {
			return models.UserInput{}, errors.New("email already exists")
		}
		return models.UserInput{}, er
	}

	return user, nil
}

func (tm *Taskmanager) Login(user models.UserInput) (string, error) {
	coll := tm.collection("users")
	var userDB models.UserInput
	query := bson.M{"email": user.Email}

	err := coll.FindOne(context.TODO(), query).Decode(&userDB)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(user.Password))

	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.Dclaims{ID: userDB.ID, Name: userDB.Name, Email: userDB.Email, IsAdmin: userDB.IsAdmin, StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Hour * 24).Unix()}})

	jwtToken, err := token.SignedString(secretKey)

	if err != nil {
		return "", err
	}

	return jwtToken, nil
}
