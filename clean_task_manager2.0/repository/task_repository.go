package repository

import (
	"context"
	"errors"
	"task_manager/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepo struct {
	coll *mongo.Collection
}

func NewTaskRepo(db *mongo.Database, name string) *TaskRepo{
	return &TaskRepo{
		coll: db.Collection(name),
	}
}

func (TR *TaskRepo) CreateTask(task domain.Task) (domain.Task, error){
	var doc bson.M
	task.ID = primitive.NewObjectID()
	bsonModel,err := bson.Marshal(task)

	if err != nil {
		return domain.Task{}, err
	}

	err = bson.Unmarshal(bsonModel , &doc)
	if err != nil {
		return domain.Task{}, err
	}

	_ , err = TR.coll.InsertOne(context.TODO(), doc)
	if err != nil {
		return domain.Task{}, err
	}

	return task, nil
}

func (TR *TaskRepo) DeleteTaskById(id string, userId primitive.ObjectID) error{
	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": obId, "user._id" : userId}

	res, err := TR.coll.DeleteOne(context.TODO(), query)

	if err != nil{
		return err
	}

	if res.DeletedCount == 0{
		return errors.New("no document with this id exists")
	}

	return nil
}

func (TR *TaskRepo) UpdateTaskById(id string, task domain.Task) (domain.Task, error){
	obId, _ := primitive.ObjectIDFromHex(id)
	task.ID = obId
	bsonModel, err := bson.Marshal(task)
	if err != nil {
		return domain.Task{}, err
	}

	var doc bson.M
	err = bson.Unmarshal(bsonModel, &doc)
	if err != nil {
		return domain.Task{}, err
	}
	filter := bson.D{{Key: "_id", Value: obId}, {Key: "user._id", Value: task.User.ID}}
	update := bson.D{{Key: "$set", Value: doc}}

	_, err = TR.coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return domain.Task{}, err
	}

	return task, nil
}

func (TR *TaskRepo) GetAllTasks(filter bson.M) ([]domain.Task, error){
	cursor, err := TR.coll.Find(context.TODO(), filter)

	if err != nil{
		return nil, err
	}

	var tasks []domain.Task

	for cursor.Next(context.TODO()){
		task := domain.Task{}
		err := cursor.Decode(&task)

		if err != nil{
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (TR *TaskRepo) FindTaskById(id string, userId primitive.ObjectID) (domain.Task, error){
	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": obId, "user._id" : userId}
	var task domain.Task
	err := TR.coll.FindOne(context.TODO(), query).Decode(&task)
	if err != nil{
		return domain.Task{}, err
	}

	return task, nil
}

