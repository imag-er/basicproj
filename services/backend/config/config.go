package config

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/joho/godotenv"
	viper "github.com/spf13/viper"
	"os"
)

type config struct {
	App      appConfig
	Logger loggerConfig
	Database databaseConfig
	Redis    redisConfig
}

type appConfig struct {
	Host string
	Port int
	Name string
}

type loggerConfig struct {
	Level string
}

type databaseConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Name     string

	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime int
}

type redisConfig struct {
	Host     string
	Port     int
	Password string
}

var Config *config

func InitConfig() {
	godotenv.Load()

	if os.Getenv("GO_ENV") == "online" {
		hlog.Info("goenv: online")
		viper.SetConfigName("online")
	} else if os.Getenv("GO_ENV") == "dev" {
		hlog.Info("goenv: dev")
		viper.SetConfigName("dev")
	} else {
		hlog.Fatal("missing env key GO_ENV")
	}

	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		hlog.Fatalf("Error reading config file, %s", err)
	}

	Config = &config{}

	if err := viper.Unmarshal(&Config); err != nil {
		hlog.Fatalf("unable to decode into struct, %s", err)
	}

	hlog.Infof("Config loaded: %+v", Config)

}
