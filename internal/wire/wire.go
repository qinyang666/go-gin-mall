//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"go-gin-mall/internal/config"
	"go-gin-mall/internal/domain/user"
	"go-gin-mall/internal/driving/repository"
	"go-gin-mall/internal/handler"
)

// 依赖注入能够编写更易于测试的代码

func wireUserHandler(c config.DBConf, logger *logrus.Logger) *handler.UserHandler {
	panic(wire.Build(handler.NewUserHandler, user.NewUserDomainService, repository.ProviderSet))
}
