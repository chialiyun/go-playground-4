package user

//go:generate mockery -name=UserRepository
type UserRepository interface {
	GetUser()
}

type Service struct {
	repository Repository
}

func (s Service) AddNewUser() {
	u := User{
		name: "grapefruit",
	}
	s.repository.SaveUser(u)
}
