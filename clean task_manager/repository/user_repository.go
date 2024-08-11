package repository

import (
	"context"
	"errors"
	"task_manager/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepo struct {
	coll *mongo.Collection
}

func (UR *UserRepo) EnsureIndexes() error {
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"email": 1}, 
		Options: options.Index().SetUnique(true),
	}
	
	_, err := UR.coll.Indexes().CreateOne(context.TODO(), indexModel)
	return err
}

func NewUserRepo(db mongo.Database, name string) (*UserRepo, error) {
	UR := &UserRepo{
		coll : db.Collection(name),
	}

	// Ensure indexes are created
	if err := UR.EnsureIndexes(); err != nil {
		return nil, err
	}

	return UR, nil
}


func (UR *UserRepo)FindByEmail(email string) (domain.UserInput, error){
	var userDB domain.UserInput
	query := bson.M{"email": email}

	err := UR.coll.FindOne(context.TODO(), query).Decode(&userDB)
	if err != nil {
		return domain.UserInput{}, err
	}

	return userDB, nil
}

func (UR *UserRepo)FindById(id string) (domain.UserInput, error){

}
func (UR *UserRepo)FindAllUsers() ([]domain.DBUser, error){

}
func (UR *UserRepo)UpdateUserById(id string, user domain.UserInput, is_admin bool) (domain.DBUser, error){

}
func (UR *UserRepo)CreateUser(user domain.UserInput) (domain.DBUser, error){
	// Set other user properties
	user.ID = primitive.NewObjectID()
	user.IsAdmin = false

	// Insert the new user
	_, er := UR.coll.InsertOne(context.TODO(), &user)
	if er != nil {
		// Check if the error is due to a duplicate key
		if mongo.IsDuplicateKeyError(er) {
			return domain.DBUser{}, errors.New("email already exists")
		}
		return domain.DBUser{}, er
	}

	return domain.ChangeToOutput(user), nil

}
func (UR *UserRepo)DeleteUserByID(id string) error{

}
