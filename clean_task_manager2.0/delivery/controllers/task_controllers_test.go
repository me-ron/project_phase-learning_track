package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"task_manager/delivery/controllers"
	"task_manager/domain"
	"task_manager/mocks"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskHandlerTestSuite struct {
	suite.Suite
	router      *gin.Engine
	Task_UC *mocks.TaskUsecase
}

func (suite *TaskHandlerTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)

	suite.router = gin.New()
	suite.Task_UC = mocks.NewTaskUsecase(suite.T())
}

func (suite *TaskHandlerTestSuite) TestGetAllTasks() {
	filter := bson.M{}
	expectedTasks := []domain.Task{
		{Title: "Task 1"},
		{Title: "Task 2"},
	}

	suite.Task_UC.On("GetTasks", filter).Return(expectedTasks, nil)

	handler := controllers.GetAllTasks(suite.Task_UC)

	req, _ := http.NewRequest(http.MethodGet, "/tasks", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Set("filter", filter)

	handler(c)
	expectedJson, err := json.Marshal(expectedTasks)
	suite.NoError(err)

	suite.Equal(http.StatusOK, w.Code)
	suite.JSONEq(string(expectedJson), w.Body.String())

	suite.Task_UC.AssertExpectations(suite.T())
}

func (suite *TaskHandlerTestSuite) TestGetTaskById(){
	expectedTask := domain.Task{Title: "test title"}
	taskId := primitive.NewObjectID().Hex()
	userId := primitive.NewObjectID()

	suite.Task_UC.On("GetTask", taskId, userId).Return(expectedTask, nil)

	handler := controllers.GetTaskById(suite.Task_UC)

	req, _ := http.NewRequest(http.MethodGet, "/tasks"+taskId, nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = req
	c.Params = gin.Params{{Key: "id", Value: taskId}}
	c.Set("user", domain.DBUser{ID: userId})

	handler(c)

	expectedTaskJson, err := json.Marshal(expectedTask)
	suite.NoError(err)

	suite.Equal(http.StatusOK, w.Code)
	suite.JSONEq(string(expectedTaskJson), w.Body.String())

	suite.Task_UC.AssertExpectations(suite.T())
}

func (suite *TaskHandlerTestSuite) TestPostTask(){
	user := domain.DBUser{}
	task := domain.Task{Title: "testTitle", User: user}

	jsonTask, _ := json.Marshal(task)

	suite.Task_UC.On("PostTask", task, user).Return(task, nil)

	handler := controllers.PostTask(suite.Task_UC)

	req, _ := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(jsonTask))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = req
	c.Set("user", user)

	handler(c)

	suite.Equal(http.StatusCreated, w.Code)
	suite.JSONEq(`{"message" : "created successfully", "task" : `+string(jsonTask)+`}`, w.Body.String())

}

func (suite *TaskHandlerTestSuite) TestDeleteTask(){
	taskId := primitive.NewObjectID().Hex()
	user := domain.DBUser{ID: primitive.NewObjectID()}

	suite.Task_UC.On("DeleteTask", taskId, user.ID).Return(nil)

	handler := controllers.DeleteTask(suite.Task_UC)

	req, _ := http.NewRequest(http.MethodDelete, "/tasks"+taskId, nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = req
	c.Params = gin.Params{{Key: "id", Value: taskId}}
	c.Set("user", user)

	handler(c)

	suite.Equal(http.StatusOK, w.Code)
	suite.JSONEq(`{"messages" : "deleted successfully"}`, w.Body.String())

	suite.Task_UC.AssertExpectations(suite.T())
}

func (suite *TaskHandlerTestSuite) TestUpdateTask(){
	user := domain.DBUser{}
	taskID := primitive.NewObjectID().Hex()
	task := domain.Task{Title: "UpdateTitle"}
	expectedTask := domain.Task{Title: "UpdateTitle", User: user}

	jsonTask, _ := json.Marshal(task)
	expectedTaskJson, _ := json.Marshal(expectedTask)

	suite.Task_UC.On("UpdateTask", taskID, task, user).Return(expectedTask, nil)

	handler := controllers.UpdateTask(suite.Task_UC)

	req, _ := http.NewRequest(http.MethodPut, "/tasks"+taskID, bytes.NewBuffer(jsonTask))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = req
	c.Params = gin.Params{{Key: "id", Value: taskID}}
	c.Set("user", user)

	handler(c)

	suite.Equal(http.StatusOK, w.Code)
	suite.JSONEq(string(expectedTaskJson), w.Body.String())

}
func TestTaskHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(TaskHandlerTestSuite))
}
