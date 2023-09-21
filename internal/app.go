package app

import (
	"context"

	"github.com/flytrap/go-template/internal/config"
	"github.com/flytrap/go-template/pkg/redis"
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
)

type options struct {
	ConfigFile string
	Version    string
}

type Option func(*options)

func SetConfigFile(s string) Option {
	return func(o *options) {
		o.ConfigFile = s
	}
}

func SetVersion(s string) Option {
	return func(o *options) {
		o.Version = s
	}
}

func InitStore() (*redis.Store, error) {
	c := redis.Config{}
	copier.Copy(&c, config.C.Redis)
	store := redis.NewStore(&c)
	return store, nil
}

func RunApp(ctx context.Context, opts ...Option) error {
	initLogger()
	var o options
	for _, opt := range opts {
		opt(&o)
	}
	config.MustLoad(o.ConfigFile)
	config.PrintWithJSON()

	injector, err := BuildInjector()
	if err != nil {
		return err
	}

	return injector.GrpcServer.Run()
}

func initLogger() {
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceQuote:      true,                  //键值对加引号
		TimestampFormat: "2006-01-02 15:04:05", //时间格式
		FullTimestamp:   true,
	})
}
