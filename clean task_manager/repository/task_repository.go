package repository

import (
	"task_manager/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepo struct {
	coll mongo.Collection
}

func NewTaskRepository(db mongo.Database, name string) *TaskRepo{
	return &TaskRepo{
		coll: *db.Collection(name),
	}
}

func (TR *TaskRepo) CreateTask(task domain.Task) (domain.Task, error){

}

func (TR *TaskRepo) DeleteTaskById(id string, userId primitive.ObjectID) error{

}

func (TR *TaskRepo) UpdateTaskById(id string, task domain.Task) (domain.Task, error){

}

func (TR *TaskRepo) GetAllTasks(filter bson.M) ([]domain.Task, error){

}

func (TR *TaskRepo) FindTaskById(id string, userId primitive.ObjectID) (domain.Task, error){

}

