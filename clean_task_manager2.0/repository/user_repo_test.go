package repository

import (
	"context"
	"testing"

	"task_manager/domain"
	"task_manager/mocks"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepoTestSuite struct {
	suite.Suite
	mockColl       *mocks.CollectionInterface
	mockCursor     *mocks.CursorInterface
	mockSingleResult *mocks.SingleResultInterface
	mockDeleteResult *mocks.DeleteResultInterface
	mockIdx 		*mocks.IndexView
	repo           *UserRepo
}

func (suite *UserRepoTestSuite) SetupTest() {
	suite.mockColl = mocks.NewCollectionInterface(suite.T())
	suite.mockCursor = mocks.NewCursorInterface(suite.T())
	suite.mockSingleResult = mocks.NewSingleResultInterface(suite.T())
	suite.mockDeleteResult = mocks.NewDeleteResultInterface(suite.T())
	suite.mockIdx = mocks.NewIndexView(suite.T())

	var err error
	suite.mockColl.On("Indexes").Return(suite.mockIdx)
	suite.mockIdx.On("CreateOne", context.TODO(), mongo.IndexModel{
		Keys:    bson.M{"email": 1}, 
		Options: options.Index().SetUnique(true),
	}).Return("", nil)

	suite.repo, err = NewUserRepo(suite.mockColl)
	suite.Require().NoError(err)
}

func (suite *UserRepoTestSuite) TestFindByEmail() {
	suite.mockColl.On("FindOne", context.TODO(), bson.M{"email": "test@example.com"}).Return(suite.mockSingleResult)
	suite.mockSingleResult.On("Decode", mock.Anything).Run(func(args mock.Arguments) {
		user := args.Get(0).(*domain.UserInput)
		*user = domain.UserInput{Email: "test@example.com"}
	}).Return(nil)

	user, err := suite.repo.FindByEmail("test@example.com")
	suite.NoError(err)
	suite.Equal("test@example.com", user.Email)

	suite.mockColl.AssertExpectations(suite.T())
	suite.mockSingleResult.AssertExpectations(suite.T())
}

func (suite *UserRepoTestSuite) TestFindById() {
	suite.mockColl.On("FindOne", context.TODO(), mock.MatchedBy(func(m interface{}) bool {
        _, ok := m.(primitive.M)
        return ok
    })).Return(suite.mockSingleResult)
	suite.mockSingleResult.On("Decode", mock.Anything).Run(func(args mock.Arguments) {
		user := args.Get(0).(*domain.UserInput)
		*user = domain.UserInput{ID: primitive.NewObjectID()}
	}).Return(nil)

	id := primitive.NewObjectID().Hex()
	user, err := suite.repo.FindById(id)
	suite.NoError(err)
	suite.NotNil(user.ID)

	suite.mockColl.AssertExpectations(suite.T())
	suite.mockSingleResult.AssertExpectations(suite.T())
}

func (suite *UserRepoTestSuite) TestFindAllUsers() {
	suite.mockColl.On("Find", context.TODO(), bson.M{}).Return(suite.mockCursor, nil)
	suite.mockCursor.On("Next", context.TODO()).Return(true).Once()
	suite.mockCursor.On("Decode", mock.Anything).Run(func(args mock.Arguments) {
		user := args.Get(0).(*domain.UserInput)
		*user = domain.UserInput{Email: "test@example.com"}
	}).Return(nil)
	suite.mockCursor.On("Next", context.TODO()).Return(false).Once()

	users, err := suite.repo.FindAllUsers()
	suite.NoError(err)
	suite.Len(users, 1)
	suite.Equal("test@example.com", users[0].Email)

	suite.mockColl.AssertExpectations(suite.T())
	suite.mockCursor.AssertExpectations(suite.T())
}

func (suite *UserRepoTestSuite) TestUpdateUserById() {
	suite.mockColl.On("UpdateOne", context.TODO(),  mock.MatchedBy(func(m interface{}) bool {
        _, ok := m.(primitive.D)
        return ok
    }),  mock.MatchedBy(func(m interface{}) bool {
        _, ok := m.(primitive.D)
        return ok
    })).Return(&mongo.UpdateResult{}, nil)

	id := primitive.NewObjectID().Hex()
	user := domain.UserInput{ID: primitive.NewObjectID(), Email: "updated@example.com"}
	updatedUser, err := suite.repo.UpdateUserById(id, user, false)
	suite.NoError(err)
	suite.Equal("updated@example.com", updatedUser.Email)

	suite.mockColl.AssertExpectations(suite.T())
}

func (suite *UserRepoTestSuite) TestCreateUser() {
	suite.mockColl.On("InsertOne", context.TODO(), mock.Anything).Return(&mongo.InsertOneResult{}, nil)

	user := domain.UserInput{Email: "new@example.com"}
	createdUser, err := suite.repo.CreateUser(user)
	suite.NoError(err)
	suite.Equal("new@example.com", createdUser.Email)

	suite.mockColl.AssertExpectations(suite.T())
}

func (suite *UserRepoTestSuite) TestDeleteUserByID() {
	suite.mockColl.On("DeleteOne", context.TODO(), mock.MatchedBy(func(m interface{}) bool {
        _, ok := m.(primitive.M)
        return ok
    })).Return(suite.mockDeleteResult, nil)
	suite.mockDeleteResult.On("DeletedCount").Return(int64(1))

	id := primitive.NewObjectID().Hex()
	err := suite.repo.DeleteUserByID(id)
	suite.NoError(err)

	suite.mockColl.AssertExpectations(suite.T())
	suite.mockDeleteResult.AssertExpectations(suite.T())
}

func TestUserRepoTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepoTestSuite))
}
