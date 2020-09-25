package redisx

import (
	"github.com/go-redis/redis"
)

const (
	defaultMode    = 0 //使用 isTencentCloudRedis 传递参数
	clusterMode    = 1
	singleNodeMode = 2
)

type Option struct {
	ClusterOptions *redis.ClusterOptions
	Options        *redis.Options
	IsCluster      bool
}

func CreateRedisClient(option *Option) (redis.Cmdable, error) {
	var ret redis.Cmdable
	if option.IsCluster {
		ret = redis.NewClusterClient(option.ClusterOptions)
	} else {
		ret = redis.NewClient(option.Options)
	}

	_, err := ret.Ping().Result()
	return ret, err
}
