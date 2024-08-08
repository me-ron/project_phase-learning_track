package data

import (
	"context"
	"errors"
	"task_manager/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
type Taskmanager struct {
	Client *mongo.Client
}

func (tm *Taskmanager) collection(name string) *mongo.Collection{
	return tm.Client.Database("task_manager").Collection(name)
}

func (tm *Taskmanager) EnsureIndexes() error {
	coll := tm.collection("users")
	
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"email": 1}, 
		Options: options.Index().SetUnique(true),
	}
	
	_, err := coll.Indexes().CreateOne(context.TODO(), indexModel)
	return err
}

func NewTaskmanager(client *mongo.Client) (*Taskmanager, error) {
	tm := &Taskmanager{
		Client: client,
	}

	// Ensure indexes are created
	if err := tm.EnsureIndexes(); err != nil {
		return nil, err
	}

	return tm, nil
}

func (tm *Taskmanager) PostTask(task models.Task) error {
	coll := tm.collection("tasks")

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

func (tm *Taskmanager) GetTasks(filter bson.M) ([]models.Task, error) {
	coll := tm.collection("tasks")
	cursor, err := coll.Find(context.TODO(), filter)

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

func (tm *Taskmanager) GetTask(id string, userId primitive.ObjectID) (models.Task, error) {
	coll := tm.collection("tasks")
	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": obId, "user._id" : userId}
	var task models.Task
	err := coll.FindOne(context.TODO(), query).Decode(&task)
	if err != nil{
		return models.Task{}, err
	}

	return task, nil

	
}

func (tm *Taskmanager) DeleteTask(id string, userId primitive.ObjectID) error {
	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": obId, "user._id" : userId}

	coll := tm.collection("tasks")
	res, err := coll.DeleteOne(context.TODO(), query)

	if err != nil{
		return err
	}

	if res.DeletedCount == 0{
		return errors.New("no document with this id exists")
	}

	return nil
	
}

func (tm *Taskmanager) UpdateTask(id string, task models.Task, user models.DBUser) (models.Task, error) {
	coll := tm.collection("tasks")
	obId, _ := primitive.ObjectIDFromHex(id)
	task.ID = obId
	task.User = user
	bsonModel, err := bson.Marshal(task)
	if err != nil {
		return models.Task{}, err
	}

	var doc bson.M
	err = bson.Unmarshal(bsonModel, &doc)
	if err != nil {
		return models.Task{}, err
	}
	filter := bson.D{{Key: "_id", Value: obId}, {Key: "user._id", Value: user.ID}}
	update := bson.D{{Key: "$set", Value: doc}}

	_, err = coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return models.Task{}, err
	}

	return task, nil
	
}