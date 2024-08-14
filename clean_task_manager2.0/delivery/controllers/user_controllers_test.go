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

	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandlerTestSuite struct {
	suite.Suite
	router      *gin.Engine
	User_UC *mocks.UserUsecase
}

func (suite *UserHandlerTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)

	suite.router = gin.New()
	suite.User_UC = mocks.NewUserUsecase(suite.T())
}

func (suite *UserHandlerTestSuite) TestRegister(){
	user := domain.UserInput{}
	usr := domain.DBUser{}

	suite.User_UC.On("Signup", user).Return(usr, nil)

	handler := controllers.Register(suite.User_UC)
	jsonuser, _ := json.Marshal(user)

	req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(jsonuser))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = req

	handler(c)
	jsonUsr, err := json.Marshal(usr)
	suite.NoError(err)

	suite.Equal(http.StatusCreated, w.Code)
	suite.JSONEq(`{"message" : "User created Successfully.", "user" : `+string(jsonUsr)+`}`, w.Body.String())

}

func (suite *UserHandlerTestSuite) TestLogIn(){
	user := domain.UserInput{}
	usr := domain.DBUser{}

	suite.User_UC.On("Login", user).Return(usr,"TOKEN", nil)

	handler := controllers.Login(suite.User_UC)
	jsonuser, _ := json.Marshal(user)

	req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(jsonuser))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = req

	handler(c)
	jsonUsr, err := json.Marshal(usr)
	suite.NoError(err)

	suite.Equal(http.StatusOK, w.Code)
	suite.JSONEq(`{"token" : "TOKEN", "user" : `+string(jsonUsr)+`}`, w.Body.String())

}

func (suite *UserHandlerTestSuite) TestGetAllTasks() {
	expectedUsers := []domain.DBUser{
		{Email: "test1@email.com"},
		{Email: "test2@email.com"},
	}

	suite.User_UC.On("GetUsers").Return(expectedUsers, nil)

	handler := controllers.GetAllUsers(suite.User_UC)

	req, _ := http.NewRequest(http.MethodGet, "/users", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	handler(c)
	expectedJson, err := json.Marshal(expectedUsers)
	suite.NoError(err)

	suite.Equal(http.StatusOK, w.Code)
	suite.JSONEq(string(expectedJson), w.Body.String())

	suite.User_UC.AssertExpectations(suite.T())
}

func (suite *UserHandlerTestSuite) TestGetUserById(){
	userId := primitive.NewObjectID().Hex()
	expectedUser := domain.DBUser{}

	suite.User_UC.On("GetUser", userId).Return(expectedUser, nil)

	handler := controllers.GetUserById(suite.User_UC)

	req, _ := http.NewRequest(http.MethodGet, "/users"+userId, nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = req
	c.Params = gin.Params{{Key: "id", Value: userId}}

	handler(c)

	expectedUserJson, err := json.Marshal(expectedUser)
	suite.NoError(err)

	suite.Equal(http.StatusAccepted, w.Code)
	suite.JSONEq(string(expectedUserJson), w.Body.String())

	suite.User_UC.AssertExpectations(suite.T())
}

func (suite *UserHandlerTestSuite) TestMakeAdmin(){
	userId := primitive.NewObjectID().Hex()
	expectedUser := domain.DBUser{IsAdmin: true}

	suite.User_UC.On("MakeAdmin", userId).Return(expectedUser, nil)

	handler := controllers.MakeAdmin(suite.User_UC)

	req, _ := http.NewRequest(http.MethodPut, "/users"+userId, nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = req
	c.Params = gin.Params{{Key: "id", Value: userId}}

	handler(c)

	expectedUserJson, err := json.Marshal(expectedUser)
	suite.NoError(err)

	suite.Equal(http.StatusAccepted, w.Code)
	suite.JSONEq(string(expectedUserJson), w.Body.String())

	suite.User_UC.AssertExpectations(suite.T())
}

func (suite *UserHandlerTestSuite) TestDeleteUser(){
	userID := primitive.NewObjectID().Hex()

	suite.User_UC.On("DeleteUser", userID).Return(nil)

	handler := controllers.DeleteUser(suite.User_UC)

	req, _ := http.NewRequest(http.MethodDelete, "/users"+userID, nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = req
	c.Params = gin.Params{{Key: "id", Value: userID}}

	handler(c)

	suite.Equal(http.StatusAccepted, w.Code)
	suite.JSONEq(`{"message" : "User deleted successfully"}`, w.Body.String())

	suite.User_UC.AssertExpectations(suite.T())
}

func (suite *UserHandlerTestSuite) TestUpdateUser(){
	user := domain.UserInput{}
	usr := domain.DBUser{}
	userId := primitive.NewObjectID().Hex()

	jsonUser, _ := json.Marshal(user)
	expectedUserJson, _ := json.Marshal(usr)

	suite.User_UC.On("UpdateUser",userId, user).Return(usr, nil)

	handler := controllers.UpdateUser(suite.User_UC)

	req, _ := http.NewRequest(http.MethodPut, "/users"+userId, bytes.NewBuffer(jsonUser))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = req
	c.Params = gin.Params{{Key: "id", Value: userId}}
	c.Set("user", user)

	handler(c)

	suite.Equal(http.StatusAccepted, w.Code)
	suite.JSONEq(string(expectedUserJson), w.Body.String())

}

func TestUserHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(UserHandlerTestSuite))
}