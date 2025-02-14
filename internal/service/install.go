package service

import (
	"CipherX/config"
	"CipherX/config/autoload"
	"CipherX/initialize"
	"CipherX/internal/core"
	"CipherX/internal/fields"
	res "CipherX/pkg/response"
	"fmt"
)

func Install(field fields.InstallFields) (res.ResCode, any) {
	// Save the configuration
	if err := core.SaveConfig(field); err != nil {
		return res.CodeGenericError, err
	}

	// Reload the configuration
	config.GinVP = initialize.Viper("config.yaml")

	// Reload the database configuration
	config.GinDB = initialize.DB()
	if config.GinDB == nil {
		fmt.Println("Failed to initialize database...")
	} else {
		// Migrating the database
		initialize.RegisterTables(config.GinDB)
	}
	config.GinRedis = initialize.Redis()
	if config.GinRedis == nil {
		fmt.Println("Failed to initialize redis...")
	}

	return res.CodeSuccess, "Configuration file generated successfully"
}

func InstallDBTest(field fields.DBTestFields) (res.ResCode, any) {
	var dbConfig autoload.DB
	dbConfig.Host = field.Host
	dbConfig.Port = field.Port
	dbConfig.Username = field.User
	dbConfig.Password = field.Pass
	dbConfig.Dbname = field.Name
	dbConfig.Config = "charset=utf8mb4&parseTime=True&loc=Local"

	b, err := core.DBTest(dbConfig.Dsn())
	if err != nil {
		return res.CodeGenericError, err
	}
	if !b {
		return res.CodeGenericError, "Database connection failed"
	}

	return res.CodeSuccess, "Database connection successful"
}

func InstallRedisTest(field fields.RedisTestFields) (res.ResCode, any) {
	b, err := core.RedisTest(field.Host, field.Pass, field.Port)
	if err != nil {
		return res.CodeGenericError, err
	}
	if !b {
		return res.CodeGenericError, "Redis connection failed"
	}

	return res.CodeSuccess, "Redis connection successful"
}
