package controller

import (
	"go/go-backend-api/internal/service"
	"go/go-backend-api/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) Register(c *gin.Context) {
	result := uc.userService.Register("", "")

	response.SuccessResponse(c, result, nil)
}

// func (uc *UserController) GetUserByID(c *gin.Context) {
// 	userInfo := uc.userService.GetUserInfo()
// 	response.SuccessResponse(c, response.ErrCodeSuccess, userInfo)
// }

// Interface version
