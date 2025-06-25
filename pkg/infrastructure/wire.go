//go:build wireinject
// +build wireinject

package infrastructure

import (
	"github.com/google/wire"
)

func InitInfra() (*RedisClient, error) {
	wire.Build(NewRedis)
	return &RedisClient{}, nil
}
