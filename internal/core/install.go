package core

import (
	"CipherX/config"
	"CipherX/internal/fields"
	"database/sql"
	"github.com/spf13/viper"
	"strconv"

	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DBTest Testing database connectivity
func DBTest(dsn string) (bool, error) {
	newDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return false, err
	}

	db, err := newDB.DB()
	if err != nil {
		return false, err
	}

	// Close connection
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			config.GinLOG.Error(err.Error())
		}
	}(db)

	err = db.Ping()
	if err != nil {
		return false, err
	}

	return true, nil
}

// RedisTest Testing redis connectivity
func RedisTest(redisAddr, redisPwd string, redisPort int) (bool, error) {
	redisDB := redis.NewClient(&redis.Options{
		Addr:     redisAddr + ":" + strconv.Itoa(redisPort),
		Password: redisPwd, // no password (default)
		DB:       0,        // use default DB
	})

	_, err := redisDB.Ping().Result()
	if err != nil {
		return false, err
	}

	// Close connection
	defer func(redisDB *redis.Client) {
		err = redisDB.Close()
		if err != nil {
			config.GinLOG.Error(err.Error())
		}
	}(redisDB)

	return true, nil
}

// SaveConfig Save configuration file
func SaveConfig(field fields.InstallFields) error {
	viper.Set("db.host", field.DB.Host)
	viper.Set("db.port", field.DB.Port)
	viper.Set("db.config", "charset=utf8mb4&parseTime=True&loc=Local")
	viper.Set("db.db-name", field.DB.Name)
	viper.Set("db.username", field.DB.User)
	viper.Set("db.password", field.DB.Pass)
	viper.Set("db.prefix", "")
	viper.Set("db.singular", false)
	viper.Set("db.engine", "")
	viper.Set("db.max-idle-conns", 10)
	viper.Set("db.max-open-conns", 100)
	viper.Set("db.log-mode", true)
	viper.Set("db.log-zap", true)

	viper.Set("redis.host", field.Redis.Host)
	viper.Set("redis.port", field.Redis.Port)
	viper.Set("redis.password", field.Redis.Pass)
	viper.Set("redis.db", 0)
	viper.Set("redis.pool-size", 100)
	return viper.SafeWriteConfigAs("config.yaml")
}
