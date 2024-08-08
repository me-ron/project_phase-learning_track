package data

// import (
// 	"context"
// 	"task_manager/models"

// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// func (tm *Taskmanager) GetUsers() ([]models.User, error){
// 	coll := tm.collection("users")
// 	query := bson.D{}
// 	cursor, err := coll.Find(context.TODO(),query)

// 	if err != nil{
// 		return nil, err
// 	}

// 	var users []models.User

// 	for cursor.Next(context.TODO()){
// 		user := models.User{}
// 		err := cursor.Decode(&user)
// 		if err != nil{
// 			return nil, err
// 		}

// 		users = append(users, user)
// 	}

// 	return users, nil
// }


// func (tm *Taskmanager) GetUser(id string) (models.User, error){
// 	coll := tm.collection("users")
// 	obId, _ := primitive.ObjectIDFromHex(id)
// }