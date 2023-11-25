package configs

import "github.com/spf13/viper"

func InitConfigs() error {
	viper.SetConfigFile("configs/config.yaml")
	return viper.ReadInConfig()
}
