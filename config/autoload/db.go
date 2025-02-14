package autoload

type GeneralDB struct {
	Host         string `mapstructure:"host" json:"host" yaml:"host"`                               // Database address
	Port         string `mapstructure:"port" json:"port" yaml:"port"`                               // Database port
	Config       string `mapstructure:"config" json:"config" yaml:"config"`                         // Advanced configuration
	Dbname       string `mapstructure:"db-name" json:"db-name" yaml:"db-name"`                      // Database name
	Username     string `mapstructure:"username" json:"username" yaml:"username"`                   // Database username
	Password     string `mapstructure:"password" json:"password" yaml:"password"`                   // Database password
	Prefix       string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                         // Global table prefix, not effective if TableName is defined separately
	Singular     bool   `mapstructure:"singular" json:"singular" yaml:"singular"`                   // Whether to enable global disabling of plural, true means enable
	Engine       string `mapstructure:"engine" json:"engine" yaml:"engine" default:"InnoDB"`        // Database engine, default InnoDB
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"` // Maximum number of idle connections
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"` // Maximum number of open connections to the database
	LogMode      string `mapstructure:"log-mode" json:"log-mode" yaml:"log-mode"`                   // Whether to enable Gorm global logging
	LogZap       bool   `mapstructure:"log-zap" json:"log-zap" yaml:"log-zap"`                      // Whether to write logs to files through zap
}

type DB struct {
	GeneralDB `yaml:",inline" mapstructure:",squash"`
}

func (m *DB) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
}

func (m *DB) GetLogMode() string {
	return m.LogMode
}
