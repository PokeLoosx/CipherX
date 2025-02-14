package initialize

import (
	"CipherX/internal/model"
	"os"

	"gorm.io/gorm"
)

func DB() *gorm.DB {
	return GormMysql()
}

// RegisterTables Register database tables exclusively
func RegisterTables(db *gorm.DB) {
	// Data tables: Automatic migration
	err := db.Set("gorm:table_options", "CHARSET=utf8mb4").AutoMigrate(
		&model.User{},
	)
	if err != nil {
		os.Exit(0)
	}
}
