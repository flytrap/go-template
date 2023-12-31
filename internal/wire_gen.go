// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"github.com/flytrap/go-template/internal/api"
	"github.com/flytrap/go-template/internal/repositories"
	"github.com/flytrap/go-template/internal/services"
)

// Injectors from wire.go:

func BuildInjector() (*Injector, error) {
	db, err := InitGormDB()
	if err != nil {
		return nil, err
	}
	logRepository := repositories.NewLoggingRepository(db)
	logService := services.NewLogService(logRepository)
	logServiceServer := api.NewLogApi(logService)
	grpcServer := InitGrpcServer(logServiceServer)
	injector := &Injector{
		GrpcServer: grpcServer,
	}
	return injector, nil
}
