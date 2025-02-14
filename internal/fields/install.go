package fields

type InstallFields struct {
	DB    DBTestFields    `json:"db" required:"true"`
	Redis RedisTestFields `json:"redis" required:"true"`
}

type DBTestFields struct {
	Host string `json:"db_host" required:"true"`
	Port string `json:"db_port" required:"true"`
	User string `json:"db_user" required:"true"`
	Pass string `json:"db_pass" required:"true"`
	Name string `json:"db_name" required:"true"`
}

type RedisTestFields struct {
	Host string `json:"redis_host" required:"true"`
	Port int    `json:"redis_port" required:"true"`
	Pass string `json:"redis_pass" required:"false"`
}
