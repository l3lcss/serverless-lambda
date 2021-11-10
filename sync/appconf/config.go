package appconf

import (
	"strings"

	"github.com/spf13/viper"
)

// MySQL struct
type MySQL struct {
	Host     string
	Port     int64
	Username string
	Password string
	Database string
}

// MSSQL struct
type MSSQL struct {
	Host     string
	Port     int64
	Username string
	Password string
	Database string
}

type app struct {
	SkipMode bool
}

// Config struct
type Config struct {
	App   *app
	Mysql *MySQL
	MSSQL *MSSQL
}

// Read func
func Read() *Config {
	conf := &Config{}
	viper.SetConfigFile(`appconf/config.yaml`)
	viper.AutomaticEnv()

	// when replacing from env vars, ex: mysql.host to MYSQL_HOST
	viper.SetEnvKeyReplacer(strings.NewReplacer(`.`, `_`))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(conf)
	if err != nil {
		panic(err)
	}

	return conf
}
