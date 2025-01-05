package user

import (
	"go/go-backend-api/internal/controller"
	"go/go-backend-api/internal/repo"
	"go/go-backend-api/internal/service"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// public router
	// this is non-dependency
	ur := repo.NewUserRepository()
	us := service.NewUserService(ur)
	userHandlerNonDependency := controller.NewUserController(us)

	// WIRE go

	// Dependency Injection (DI)

	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/register", userHandlerNonDependency.Register) // register -> YES -> NO
		userRouterPublic.POST("/otp")
	}

	// private router
	userRouterPrivate := Router.Group("/user")
	// userRouterPrivate.User(limiter())
	// userRouterPrivate.Use(Auth())
	// userRouterPrivate.Use(Permission())
	{
		userRouterPrivate.GET("/get_info")
	}

}
