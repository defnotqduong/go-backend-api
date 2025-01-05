package service

import (
	"go/go-backend-api/internal/repo"
	"go/go-backend-api/pkg/response"
)

// type UserService struct {
// 	userRepo *repo.UserRepo
// }

// func NewUserService() *UserService {
// 	return &UserService{
// 		userRepo: repo.NewUserRepo(),
// 	}
// }

// func (us *UserService) GetUserInfo() string {
// 	return us.userRepo.GetInfoUser()
// }

// Interface version

type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	userRepo repo.IUserRepository
}

// Register implements IUserService.
func (us *userService) Register(email string, purpose string) int {
	// 1. check email exists

	if us.userRepo.GetUserByEmail(email) {
		return response.ErrCodeUserHasExists
	}

	return response.ErrCodeSuccess

}

func NewUserService(userRepo repo.IUserRepository) IUserService {
	return &userService{
		userRepo: userRepo,
	}
}
