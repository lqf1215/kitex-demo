package config

type ClientConfig struct {
	DB     *DBConfig    `toml:"db"`
	Server ServerConfig `toml:"server"'`
}

type ServerConfig struct {
	Port int64 `toml:"port"'`
}

type DBConfig struct {
	dsn          string `mapstructure:"dsn" json:"username" yaml:"username"`                      // 数据库用户名
	MaxIdleConns int    `mapstructure:"maxIdleConns" json:"max-idle-conns" yaml:"max-idle-conns"` // 空闲中的最大连接数
	MaxOpenConns int    `mapstructure:"maxOpenConns" json:"max-open-conns" yaml:"max-open-conns"` // 打开到数据库的最大连接数
}

func DefaultClientConfig() ClientConfig {
	return ClientConfig{
		DB: &DBConfig{
			dsn:          "root:root@tcp(127.0.0.1:3306)/kitex_demo?charset=utf8&parseTime=True&loc=Local",
			MaxIdleConns: 5,
			MaxOpenConns: 20,
		},
		Server: ServerConfig{Port: 8888},
	}
}
