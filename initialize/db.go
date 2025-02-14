package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"CipherX/config"
)

type Writer struct {
	logger.Writer
}

// NewWriter Writer constructor
func NewWriter(w logger.Writer) *Writer {
	return &Writer{Writer: w}
}

// Printf Format and print logs
func (w *Writer) Printf(message string, data ...interface{}) {
	var logZap bool
	logZap = config.GinConfig.DB.LogZap
	if logZap {
		config.GinLOG.Info(fmt.Sprintf(message+"\n", data...))
	} else {
		w.Writer.Printf(message, data...)
	}
}

type DbBase interface {
	GetLogMode() string
}

var orm = new(_gorm)

type _gorm struct{}

// Config gorm custom configuration
func (g *_gorm) Config(prefix string, singular bool) *gorm.Config {
	cfg := &gorm.Config{
		// Naming strategy
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   prefix,   // Table prefix, add prefix before table name, e.g., add user module table prefix 'user_'
			SingularTable: singular, // Whether to use singular form for table names, if set to true, User model will correspond to users table
		},

		DisableForeignKeyConstraintWhenMigrating: true,
	}
	_default := logger.New(NewWriter(log.New(os.Stdout, "\r\n", log.LstdFlags)), logger.Config{
		SlowThreshold: 200 * time.Millisecond,
		LogLevel:      logger.Warn,
		Colorful:      true,
	})

	var logMode DbBase
	logMode = &config.GinConfig.DB

	switch logMode.GetLogMode() {
	case "silent", "Silent":
		cfg.Logger = _default.LogMode(logger.Silent)
	case "error", "Message":
		cfg.Logger = _default.LogMode(logger.Error)
	case "warn", "Warn":
		cfg.Logger = _default.LogMode(logger.Warn)
	case "info", "Info":
		cfg.Logger = _default.LogMode(logger.Info)
	default:
		cfg.Logger = _default.LogMode(logger.Info)
	}
	return cfg

}

func GormMysql() *gorm.DB {
	m := config.GinConfig.DB
	if m.Dbname == "" {
		return nil
	}
	if gin.Mode() == gin.DebugMode {
		fmt.Println(m.Dsn())
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // Default length for string type fields
		SkipInitializeWithVersion: false,   // Auto-configure based on version
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), orm.Config(m.Prefix, m.Singular)); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}
