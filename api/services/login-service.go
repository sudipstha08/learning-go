package service

type LoginService interface {
	Login(username string, password string) bool
}

type loginService struct {
	authorizedUsername string
	authorizedPassword string
}

func NewLoginService() LoginService {
	return &loginService{
		authorizedUsername: "sudipstha",
		authorizedPassword: "pass1234",
	}
}

func (service *loginService) Login(username string, password string) bool {
	return service.authorizedUsername == username && service.authorizedPassword == password
}
