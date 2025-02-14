package initialize

import (
	"strconv"

	"github.com/go-redis/redis"

	"CipherX/config"
)

// Redis initializes Redis
func Redis() *redis.Client {
	r := config.GinConfig.Redis
	RDB := redis.NewClient(&redis.Options{
		Addr:     r.Host + ":" + strconv.Itoa(r.Port),
		Password: r.Password, // no password set
		DB:       r.DB,       // use default DB
		PoolSize: r.PoolSize, // connection pool size
	})

	_, err := RDB.Ping().Result()
	if err != nil {
		panic(err)
	}
	return RDB
}
