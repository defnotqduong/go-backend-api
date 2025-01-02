package global

import (
	"go/go-backend-api/pkg/logger"
	"go/go-backend-api/pkg/setting"

	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
)

/*
Config
Redis
Mysql
*/
