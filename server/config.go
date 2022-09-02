package main

import "github.com/spf13/viper"

type Config struct {
	Shard1 string `mapstructure:"SHARD1"`
	Shard2 string `mapstructure:"SHARD2"`
	Shard3 string `mapstructure:"SHARD3"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
