package useCase

import (
	"task_manager/domain"
	"task_manager/mocks"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskUsecaseSuite struct {
	suite.Suite
	mockRepo *mocks.TaskRepository
	Task_UC *TaskUC
}

func (suite *TaskUsecaseSuite) SetupTest(){
	suite.mockRepo = mocks.NewTaskRepository(suite.T())

	suite.Task_UC = NewTaskUC(suite.mockRepo)
}

func (suite *TaskUsecaseSuite) TestPostTask(){
	expetedTask := domain.Task{Title: "test title"}
	suite.mockRepo.On("CreateTask", domain.Task{}).Return(expetedTask, nil)

	task, err := suite.Task_UC.PostTask(domain.Task{}, domain.DBUser{})

	suite.NoError(err)
	suite.Equal(expetedTask, task)

	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskUsecaseSuite) TestGetTasks() {
	filter := bson.M{"userId": "some-user-id"}
	expectedTasks := []domain.Task{{Title: "Task 1"}, {Title: "Task 2"}}

	suite.mockRepo.On("GetAllTasks", filter).Return(expectedTasks, nil)

	tasks, err := suite.Task_UC.GetTasks(filter)

	suite.NoError(err)
	suite.Equal(expectedTasks, tasks)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskUsecaseSuite) TestGetTask() {
	id := "task-id"
	userId := primitive.NewObjectID()
	expectedTask := domain.Task{Title: "Task 1"}

	suite.mockRepo.On("FindTaskById", id, userId).Return(expectedTask, nil)

	task, err := suite.Task_UC.GetTask(id, userId)

	suite.NoError(err)
	suite.Equal(expectedTask, task)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskUsecaseSuite) TestUpdateTask() {
	id := "task-id"
	user := domain.DBUser{Email: "test@email.com"}
	task := domain.Task{Title: "Updated Task"}

	suite.mockRepo.On("UpdateTaskById", id, domain.Task{Title: "Updated Task", User: user}).Return(task, nil)

	updatedTask, err := suite.Task_UC.UpdateTask(id, task, user)

	suite.NoError(err)
	suite.Equal(task, updatedTask)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskUsecaseSuite) TestDeleteTask() {
	id := "task-id"
	userId := primitive.NewObjectID()

	suite.mockRepo.On("DeleteTaskById", id, userId).Return(nil)

	err := suite.Task_UC.DeleteTask(id, userId)

	suite.NoError(err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func TestTaskUseCaseSuite(t *testing.T){
	suite.Run(t, new(TaskUsecaseSuite))
}
