package users

var (
	UserService UserServiceInterface = &userService{}
)

type UserServiceInterface interface {
	Create(user *User) (*User, error)
	GetOne(code string) (user *User, errors error)
	GetAll() ([]User, error)
	Update(code string, user *User) (*User, error)
	Delete(code string) (string, error)
}
type userService struct {
	repo UserrepoInterface
}

func NewUserService(repository UserrepoInterface) UserServiceInterface {
	return &userService{
		repository,
	}
}
func (service *userService) Create(user *User) (*User, error) {
	return service.repo.Create(user)
}

func (service *userService) GetAll() ([]User, error) {
	return service.repo.GetAll()
}
func (service *userService) GetOne(code string) (*User, error) {
	return service.repo.GetOne(code)
}
func (service *userService) Update(code string, user *User) (*User, error) {
	return service.repo.Update(code, user)
}

func (service *userService) Delete(id string) (string, error) {
	return service.repo.Delete(id)
}
