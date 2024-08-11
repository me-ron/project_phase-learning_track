package useCase

import "task_manager/domain"

type UserUC struct {
	repo domain.UserRepository
}

func NewUserUC(repository domain.UserRepository) *UserUC{
	return &UserUC{
		repo: repository,
	}
}

func (UUC *UserUC)Login(domain.UserInput) (domain.DBUser, string, error)
func (UUC *UserUC)Signup(domain.UserInput) (domain.DBUser, error)
func (UUC *UserUC)GetUsers() ([]domain.DBUser, error)
func (UUC *UserUC)GetUser(string) (domain.DBUser, error)
func (UUC *UserUC)MakeAdmin(string) (domain.DBUser, error)
func (UUC *UserUC)UpdateUser(string, domain.UserInput) (domain.DBUser, error)
func (UUC *UserUC)DeleteUser(string) error