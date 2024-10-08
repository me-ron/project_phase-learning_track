package data

import (
	"context"
	"errors"
	"task_manager/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)
type Taskmanager struct {
	Client *mongo.Client
}

func (tm *Taskmanager) collection() *mongo.Collection{
	return tm.Client.Database("task_manager").Collection("tasks")
}

func (tm *Taskmanager) PostTask(task models.Task) error {
	coll := tm.collection()

	var doc bson.M
	bsonModel,err := bson.Marshal(task)

	if err != nil {
		return err
	}

	err = bson.Unmarshal(bsonModel , &doc)
	if err != nil {
		return err
	}

	_ , err = coll.InsertOne(context.TODO(), doc)
	if err != nil {
		return err
	}

	return nil
}

func (tm *Taskmanager) GetTasks() ([]models.Task, error) {
	coll := tm.collection()
	cursor, err := coll.Find(context.TODO(), bson.M{})

	if err != nil{
		return nil, err
	}

	var tasks []models.Task

	for cursor.Next(context.TODO()){
		task := models.Task{}
		err := cursor.Decode(&task)

		if err != nil{
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
	
}

func (tm *Taskmanager) GetTask(id string) (models.Task, error) {
	coll := tm.collection()
	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": obId}

	var task models.Task
	err := coll.FindOne(context.TODO(), query).Decode(&task)
	if err != nil{
		return models.Task{}, err
	}

	return task, nil

	
}

func (tm *Taskmanager) DeleteTask(id string) error {
	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": obId}

	coll := tm.collection()
	res, err := coll.DeleteOne(context.TODO(), query)

	if err != nil{
		return err
	}

	if res.DeletedCount == 0{
		return errors.New("no document with this id exists")
	}

	return nil
	
}

func (tm *Taskmanager) UpdateTask(id string, task models.Task) (models.Task, error) {
	coll := tm.collection()
	
	bsonModel, err := bson.Marshal(task)
	if err != nil {
		return models.Task{}, err
	}

	var doc bson.M
	err = bson.Unmarshal(bsonModel, &doc)
	if err != nil {
		return models.Task{}, err
	}

	obId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: obId}}
	update := bson.D{{Key: "$set", Value: doc}}

	_, err = coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return models.Task{}, err
	}

	return task, nil
	
}