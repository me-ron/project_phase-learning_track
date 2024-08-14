package useCase_test

import (
	"task_manager/domain"
	"task_manager/mocks"
	"task_manager/useCase"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserUsecaseSuite struct {
	suite.Suite
	mockRepo *mocks.UserRepository
	mockPass_s *mocks.PasswordService
	mockToken_s *mocks.TokenService
	User_UC *useCase.UserUC
}

func (suite *UserUsecaseSuite) SetupTest(){
	suite.mockRepo = mocks.NewUserRepository(suite.T())
	suite.mockPass_s = mocks.NewPasswordService(suite.T())
	suite.mockToken_s = mocks.NewTokenService(suite.T())

	suite.User_UC = useCase.NewUserUC(suite.mockRepo, suite.mockPass_s, suite.mockToken_s)
}

func (suite *UserUsecaseSuite) TestLogin(){
	suite.mockRepo.On("FindByEmail","test@email.com").Return(domain.UserInput{Email:"test@email.com"}, nil)
	suite.mockPass_s.On("ComparePassword", "", "").Return(true, nil)
	suite.mockToken_s.On("CreateToken", mock.MatchedBy(func(m interface{})bool{
		_, ok := m.(domain.UserInput)
		return ok
	})).Return("NewToken", nil)

	user, token, err := suite.User_UC.Login(domain.UserInput{Email: "test@email.com"})
	suite.NoError(err)
	suite.Equal("NewToken", token)
	suite.Equal("test@email.com", user.Email)

	suite.mockRepo.AssertExpectations(suite.T())
	suite.mockPass_s.AssertExpectations(suite.T())
	suite.mockToken_s.AssertExpectations(suite.T())
}

func (suite *UserUsecaseSuite) TestSignup(){
	suite.mockPass_s.On("HashPasword", "").Return("Hashed password", nil)
	suite.mockRepo.On("CreateUser", domain.UserInput{Email:"test@email.com", 
										Password: "Hashed password"}).Return(domain.DBUser{Email: "test@email.com"}, nil)
	user, err := suite.User_UC.Signup(domain.UserInput{Email:"test@email.com"})
	suite.NoError(err)
	suite.Equal("test@email.com", user.Email)

	suite.mockRepo.AssertExpectations(suite.T())
	suite.mockPass_s.AssertExpectations(suite.T())
}

func (suite *UserUsecaseSuite) TestGetUsers() {
	expectedUsers := []domain.DBUser{
		{Email: "user1@email.com"},
		{Email: "user2@email.com"},
	}
	suite.mockRepo.On("FindAllUsers").Return(expectedUsers, nil)

	users, err := suite.User_UC.GetUsers()

	suite.NoError(err)
	suite.Equal(expectedUsers, users)

	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *UserUsecaseSuite) TestGetUser() {
	expectedUser := domain.DBUser{Email: "test@email.com"}
	suite.mockRepo.On("FindById", "some-id").Return(domain.UserInput{Email: "test@email.com"}, nil)

	user, err := suite.User_UC.GetUser("some-id")

	suite.NoError(err)
	suite.Equal(expectedUser, user)

	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *UserUsecaseSuite) TestMakeAdmin() {
	expectedUser := domain.DBUser{Email: "test@email.com", IsAdmin: true}
	suite.mockRepo.On("FindById", "some-id").Return(domain.UserInput{Email: "test@email.com"}, nil)
	suite.mockRepo.On("UpdateUserById", "some-id", domain.UserInput{Email: "test@email.com"}, true).Return(expectedUser, nil)

	user, err := suite.User_UC.MakeAdmin("some-id")

	suite.NoError(err)
	suite.Equal(expectedUser, user)

	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *UserUsecaseSuite) TestUpdateUser() {
	existingUser := domain.UserInput{Email: "test@email.com", Password: "oldpassword"}
	updatedUser := domain.UserInput{Email: "updated@email.com", Password: "newpassword"}
	hashedPassword := "hashedpassword"

	suite.mockRepo.On("FindById", "some-id").Return(existingUser, nil)
	suite.mockPass_s.On("HashPasword", "newpassword").Return(hashedPassword, nil)
	suite.mockRepo.On("UpdateUserById", "some-id", mock.AnythingOfType("domain.UserInput"), false).
		Return(domain.DBUser{Email: "updated@email.com"}, nil)


	user, err := suite.User_UC.UpdateUser("some-id", updatedUser)


	suite.NoError(err)
	suite.Equal("updated@email.com", user.Email)

	suite.mockRepo.AssertExpectations(suite.T())
	suite.mockPass_s.AssertExpectations(suite.T())
}

func (suite *UserUsecaseSuite) TestDeleteUser() {
	suite.mockRepo.On("DeleteUserByID", "some-id").Return(nil)

	err := suite.User_UC.DeleteUser("some-id")

	suite.NoError(err)

	suite.mockRepo.AssertExpectations(suite.T())
}


func TestUserUseCaseSuite(t *testing.T){
	suite.Run(t, new(UserUsecaseSuite))
}