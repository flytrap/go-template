//go:build wireinject
// +build wireinject

package app

import (
	"github.com/flytrap/go-template/internal/api"
	"github.com/flytrap/go-template/internal/repositories"
	"github.com/flytrap/go-template/internal/services"
	"github.com/google/wire"
)

func BuildInjector() (*Injector, error) {
	wire.Build(
		// InitStore,
		InitGormDB,
		repositories.RepositorySet,
		InitGrpcServer,
		// InitGateway,
		services.ServiceSet,
		api.APISet,
		// router.RouterSet,
		InjectorSet,
	)
	return new(Injector), nil
}
