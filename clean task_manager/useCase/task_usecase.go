package useCase

import (
	"task_manager/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskUC struct {
	repo domain.TaskRepository
}

func NewTaskUC(repository domain.TaskRepository) *TaskUC{
	return &TaskUC{
		repo: repository,
	}
}

func (TUC *TaskUC) PostTask(task domain.Task, user domain.DBUser) (domain.Task, error) {

}

func (TUC *TaskUC) GetTasks(filter bson.M) ([]domain.Task, error){

}

func (TUC *TaskUC) GetTask(id string, userId primitive.ObjectID) (domain.Task, error){

}

func (TUC *TaskUC) UpdateTask(id string, task domain.Task, user domain.DBUser) (domain.Task, error){

}

func (TUC *TaskUC) DeleteTask(id string, userId primitive.ObjectID) error{
	
}