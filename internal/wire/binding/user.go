package binding

import (
	"git.kldmp.com/learning/demo/pkg/log"
	"github.com/google/wire"
)

type User struct {
}

type UserService struct {
	userRepo UserRepository
}

func NewUserService(userRepo UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (u *UserService) UserExist(id int) bool {
	_, err := u.userRepo.GetUserByID(id)
	return err == nil
}

type UserRepository interface {
	GetUserByID(id int) (*User, error)
}

type mockUserRepo struct {
	foo string
	bar int
}

func (u *mockUserRepo) GetUserByID(id int) (*User, error) {
	return &User{}, nil
}

func NewMockUserRepo(foo string, bar int) *mockUserRepo {
	return &mockUserRepo{
		foo: foo,
		bar: bar,
	}
}

var MockUserRepoSet = wire.NewSet(NewMockUserRepo, wire.Bind(new(UserRepository), new(*mockUserRepo)))

func Demo() {
	u := InitializeUserService("foo", 0)
	log.Info("User result:%t", u.UserExist(1))
}
