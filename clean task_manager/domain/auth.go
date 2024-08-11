package domain

type TokenService interface{
	TokenValidate(string) error
	CreateToken(UserInput)(string, error)
}