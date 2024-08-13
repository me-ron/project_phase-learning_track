package repository

import (
	"context"
	"task_manager/domain"
	"task_manager/mocks"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepoTestSuite struct {
	suite.Suite
	mockColl         *mocks.CollectionInterface
	mockCursor       *mocks.CursorInterface
	mockSingleResult *mocks.SingleResultInterface
	mockDeleteResult *mocks.DeleteResultInterface
	repo             *TaskRepo
}

func (suite *TaskRepoTestSuite) SetupTest() {
	suite.mockColl = mocks.NewCollectionInterface(suite.T())
	suite.mockCursor = mocks.NewCursorInterface(suite.T())
	suite.mockSingleResult = mocks.NewSingleResultInterface(suite.T())
	suite.mockDeleteResult = mocks.NewDeleteResultInterface(suite.T())

	suite.repo = NewTaskRepo(suite.mockColl)
}

func (suite *TaskRepoTestSuite) TestFindById() {
	suite.mockColl.On("FindOne", context.TODO(), mock.MatchedBy(func(m interface{}) bool {
        _, ok := m.(primitive.M)
        return ok
    })).Return(suite.mockSingleResult)
	suite.mockSingleResult.On("Decode", mock.Anything).Run(func(args mock.Arguments) {
		task := args.Get(0).(*domain.Task)
		*task = domain.Task{ID: primitive.NewObjectID()}
	}).Return(nil)

	id := primitive.NewObjectID()
	task, err := suite.repo.FindTaskById("0", id)
	suite.NoError(err)
	suite.NotNil(task.ID)

	suite.mockColl.AssertExpectations(suite.T())
	suite.mockSingleResult.AssertExpectations(suite.T())
}

func (suite *TaskRepoTestSuite) TestGetAlltasks() {
	suite.mockColl.On("Find", context.TODO(), bson.M{}).Return(suite.mockCursor, nil)
	suite.mockCursor.On("Next", context.TODO()).Return(true).Once()
	suite.mockCursor.On("Decode", mock.Anything).Run(func(args mock.Arguments) {
		user := args.Get(0).(*domain.Task)
		*user = domain.Task{Title: "test title"}
	}).Return(nil)
	suite.mockCursor.On("Next", context.TODO()).Return(false).Once()

	tasks, err := suite.repo.GetAllTasks(bson.M{})
	suite.NoError(err)
	suite.Len(tasks, 1)
	suite.Equal("test title", tasks[0].Title)

	suite.mockColl.AssertExpectations(suite.T())
	suite.mockCursor.AssertExpectations(suite.T())
}

func (suite *TaskRepoTestSuite) TestUpdateTaskById() {
	suite.mockColl.On("UpdateOne", context.TODO(),  mock.MatchedBy(func(m interface{}) bool {
        _, ok := m.(primitive.D)
        return ok
    }),  mock.MatchedBy(func(m interface{}) bool {
        _, ok := m.(primitive.D)
        return ok
    })).Return(&mongo.UpdateResult{}, nil)

	id := primitive.NewObjectID().Hex()
	task := domain.Task{ID: primitive.NewObjectID(), Title: "updated title"}
	updatedUser, err := suite.repo.UpdateTaskById(id, task)
	suite.NoError(err)
	suite.Equal("updated title", updatedUser.Title)

	suite.mockColl.AssertExpectations(suite.T())
}

func (suite *TaskRepoTestSuite) TestCreateTask() {
	suite.mockColl.On("InsertOne", context.TODO(), mock.Anything).Return(&mongo.InsertOneResult{}, nil)

	task := domain.Task{Title: "test title"}
	createdtask, err := suite.repo.CreateTask(task)
	suite.NoError(err)
	suite.Equal("test title", createdtask.Title)

	suite.mockColl.AssertExpectations(suite.T())
}

func (suite *TaskRepoTestSuite) TestDeleteTaskByID() {
	suite.mockColl.On("DeleteOne", context.TODO(), mock.MatchedBy(func(m interface{}) bool {
        _, ok := m.(primitive.M)
        return ok
    })).Return(suite.mockDeleteResult, nil)
	suite.mockDeleteResult.On("DeletedCount").Return(int64(1))

	id := primitive.NewObjectID().Hex()
	err := suite.repo.DeleteTaskById(id, primitive.NewObjectID())
	suite.NoError(err)

	suite.mockColl.AssertExpectations(suite.T())
	suite.mockDeleteResult.AssertExpectations(suite.T())
}

func TestTaskRepoTestSuite(t *testing.T) {
	suite.Run(t, new(TaskRepoTestSuite))
}
