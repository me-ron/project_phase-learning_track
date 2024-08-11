package useCase

import "task_manager/domain"

type UserUC struct {
	repo domain.UserRepository
	PasswordS domain.PasswordService
	TokenS domain.TokenService
}

func NewUserUC(repository domain.UserRepository, PS domain.PasswordService, TS domain.TokenService) *UserUC{
	return &UserUC{
		repo: repository,
		PasswordS: PS,
		TokenS: TS,
	}
}

func (UUC *UserUC)Login(user domain.UserInput) (domain.DBUser, string, error){
	usr, err := UUC.repo.FindByEmail(user.Email)
	if err != nil {
		return domain.DBUser{}, "", err
	}

	_, er := UUC.PasswordS.ComparePassword(usr.Password, user.Password)
	if er != nil{
		return domain.DBUser{}, "", er
	}

	token, terr := UUC.TokenS.CreateToken(usr)
	if terr != nil{
		return domain.DBUser{}, "", terr
	}

	return domain.ChangeToOutput(usr), token, nil


}
func (UUC *UserUC)Signup(user domain.UserInput) (domain.DBUser, error){
	hashed_password, err := UUC.PasswordS.HashPasword(user.Password)
	if err != nil{
		return domain.DBUser{}, err
	}

	user.Password = hashed_password
	usr, er := UUC.repo.CreateUser(user)

	if er != nil{
		return domain.DBUser{}, er
	}

	return usr, nil

}
func (UUC *UserUC)GetUsers() ([]domain.DBUser, error){
	return UUC.repo.FindAllUsers()
}
func (UUC *UserUC)GetUser(id string) (domain.DBUser, error){
	user, err := UUC.repo.FindById(id)
	return domain.ChangeToOutput(user), err
}
func (UUC *UserUC)MakeAdmin(id string) (domain.DBUser, error){
	user, err := UUC.repo.FindById(id)
	if err != nil{
		return domain.DBUser{}, err
	}
	return UUC.repo.UpdateUserById(id, user, true)
}
func (UUC *UserUC)UpdateUser(id string, user domain.UserInput) (domain.DBUser, error){
	return UUC.repo.UpdateUserById(id, user, false)
}
func (UUC *UserUC)DeleteUser(id string) error{
	return UUC.repo.DeleteUserByID(id)
}