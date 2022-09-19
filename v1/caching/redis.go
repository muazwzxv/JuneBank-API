package caching

import (
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var rds = new(RedisCacheInstance)

type RedisCacheInstance struct {
	cache  *redis.Client
	config *RedisConfiguration
}

type RedisConfiguration struct {
	Host     string
	Port     int
	User     string
	Password string
}

func GetRedisInstance() *RedisCacheInstance {
	if !rds.isInstantiated() {
		config := readConfig()
		fmt.Println(config)
		rds = newRedis(config)
	}
	return rds
}

func newRedis(config *RedisConfiguration) *RedisCacheInstance {
	dsn := fmt.Sprintf("%s:%d",
		config.Host,
		config.Port,
	)

	rds := redis.NewClient(&redis.Options{
		Addr:     dsn,
		Password: config.Password,
		DB:       0,
	})

	if rds == nil {
		panic("Failed to connect to redis")
	}

	fmt.Println(rds)

	return &RedisCacheInstance{rds, config}
}

func (r *RedisCacheInstance) isInstantiated() bool {
	return r.cache != nil
}

func readConfig() *RedisConfiguration {
	reader := viper.New()
	reader.SetConfigFile("config.yaml")

	if err := reader.ReadInConfig(); err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	return &RedisConfiguration{
		Host:     reader.GetString("redist.host"),
		Port:     reader.GetInt("redis.port"),
		User:     reader.GetString("redis.user"),
		Password: reader.GetString("redis.password"),
	}
}
