package autoload

type Redis struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`                // redis host
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`                // redis port
	Password string `mapstructure:"password" json:"password" yaml:"password"`    // redis password
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`                      // redis database
	PoolSize int    `mapstructure:"pool-size" json:"pool-size" yaml:"pool-size"` // connection pool size
}
