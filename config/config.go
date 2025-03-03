package config

import "github.com/spf13/viper"


type Config struct {
	API_KEY string
}



func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return &Config{
		API_KEY: viper.GetString("API_KEY"),
	}, nil
}