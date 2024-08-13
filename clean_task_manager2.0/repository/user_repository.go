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
	coll domain.CollectionInterface
}

func (UR *UserRepo) EnsureIndexes() error {
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"email": 1}, 
		Options: options.Index().SetUnique(true),
	}
	
	_, err := UR.coll.Indexes().CreateOne(context.TODO(), indexModel)
	return err
}

func NewUserRepo(collection domain.CollectionInterface) (*UserRepo, error) {
	UR := &UserRepo{
		coll : collection,
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
	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": obId}
	var user domain.UserInput
	err := UR.coll.FindOne(context.TODO(), query).Decode(&user)
	if err != nil{
		return domain.UserInput{}, err
	}

	return user, nil
}

func (UR *UserRepo)FindAllUsers() ([]domain.DBUser, error){
	cursor, err := UR.coll.Find(context.TODO(), bson.M{})

	if err != nil{
		return nil, err
	}

	var users []domain.DBUser

	for cursor.Next(context.TODO()){
		user := domain.UserInput{}
		err := cursor.Decode(&user)

		if err != nil{
			return nil, err
		}

		users = append(users, domain.ChangeToOutput(user))
	}

	return users, nil
}

func (UR *UserRepo)UpdateUserById(id string, user domain.UserInput, is_admin bool) (domain.DBUser, error){
	obId, _ := primitive.ObjectIDFromHex(id)
	user.ID = obId
	user.IsAdmin = is_admin
	// bsonModel, err := bson.Marshal(user)
	// if err != nil {
	// 	return domain.DBUser{}, err
	// }

	// var doc bson.M
	// err = bson.Unmarshal(bsonModel, &doc)
	// if err != nil {
	// 	return domain.DBUser{}, err
	// }
	filter := bson.D{{Key: "_id", Value: obId}, {Key: "_id", Value: user.ID}}
	update := bson.D{{Key: "$set", Value: user}}

	_, err := UR.coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return domain.DBUser{}, err
	}

	return domain.ChangeToOutput(user), nil
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
	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": obId}

	res, err := UR.coll.DeleteOne(context.TODO(), query)

	if err != nil{
		return err
	}

	if res.DeletedCount() == 0{
		return errors.New("no document with this id exists")
	}

	return nil
}
