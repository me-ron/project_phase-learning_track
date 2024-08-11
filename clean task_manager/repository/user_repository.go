package repository

import (
	"context"
	"task_manager/domain"

	"go.mongodb.org/mongo-driver/bson"
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


func (UR *UserRepo)FindByEmail(string) (domain.UserInput, error)
func (UR *UserRepo)FindById(string) (domain.UserInput, error)
func (UR *UserRepo)FindAllUsers() ([]domain.DBUser, error)
func (UR *UserRepo)UpdateUserById(string, domain.UserInput, bool) (domain.DBUser, error)
func (UR *UserRepo)CreateUser(domain.UserInput) (domain.DBUser, error)
func (UR *UserRepo)DeleteUserByID(string) error
