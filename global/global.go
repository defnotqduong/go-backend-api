package global

import (
	"go/go-backend-api/pkg/logger"
	"go/go-backend-api/pkg/setting"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
)

/*
Config
Redis
Mysql
*/
