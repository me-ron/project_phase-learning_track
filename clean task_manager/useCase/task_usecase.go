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
	task.User = user
	return TUC.repo.CreateTask(task)
}

func (TUC *TaskUC) GetTasks(filter bson.M) ([]domain.Task, error){
	return TUC.repo.GetAllTasks(filter)
}

func (TUC *TaskUC) GetTask(id string, userId primitive.ObjectID) (domain.Task, error){
	return TUC.repo.FindTaskById(id, userId)
}

func (TUC *TaskUC) UpdateTask(id string, task domain.Task, user domain.DBUser) (domain.Task, error){
	task.User = user
	return TUC.repo.UpdateTaskById(id, task)
}

func (TUC *TaskUC) DeleteTask(id string, userId primitive.ObjectID) error{
	return TUC.repo.DeleteTaskById(id, userId)
}