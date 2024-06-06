package environment

import (
	"mummy-storage/config"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

type SvcEnv struct {
}

func NewSvc(rds *redis.Redis, c *config.Config) *SvcEnv {
	return &SvcEnv{}
}
