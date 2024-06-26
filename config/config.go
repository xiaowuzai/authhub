package config

import (
	"time"

	"github.com/spf13/viper"
)

// Config 存储应用的所有配置, 使用 viper 来管理
type Config struct {
	Environment       string `mapstructure:"ENVIRONMENT"`
	DBDriver          string `mapstructure:"DB_DRIVER"`
	DBSource          string `mapstructure:"DB_SOURCE"`
	MigrationUrl      string `mapstructure:"MIGRATION_URL"`
	RedisAddress      string `mapstructure:"REDIS_ADDRESS"`
	HTTPServerAddress string `mapstructure:"HTTP_SERVER_ADDRESS"`
	GRPCServerAddress string `mapstructure:"GRPC_SERVER_ADDRESS"`
	TokenSymmetricKey string `mapstructure:"TOKEN_SYMMETRIC_KEY"` //
	// EmailSenderName      string        `mapstructure:"EMAIL_SENDER_NAME"`
	// EmailSenderAddress   string        `mapstructure:"EMAIL_SENDER_ADDRESS"`
	// EmailSenderPassword  string        `mapstructure:"EMAIL_SENDER_PASSWORD"`
	// EmailVerfifyHost     string        `mapstructure:"EMAIL_VERIFY_HOST"`
	TokenDuration        time.Duration `mapstructure:"TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	AppKeyDuration       time.Duration `mapstructure:"APP_KEY_DURATION"`
	AppKey               string        `mapstructure:"APP_KEY"`
	AppKeySecret         string        `mapstructure:"APP_SECRET"`
	SiriusHost           string        `mapstructure:"SIRIUS_HOST"`
	SiriusPort           string        `mapstructure:"SIRIUS_PORT"`
}

// LoadConfig  从指定路径中加载配置
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env") //json yml

	viper.AutomaticEnv() // 用 ENV 替换默认的环境变量值

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
