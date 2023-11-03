package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Port  string `mapstructure:"PORT"`
	Store string `mapstructure:"STORE"`
}

func Load() (cfg Config) {
	v := viper.New()

	v.AddConfigPath("../configs")
	v.SetConfigName("service")
	v.SetConfigType("env")

	v.AutomaticEnv()

	var err error

	if err = v.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("Ошибка чтения файла: %s", err.Error()))
	}

	if err = v.Unmarshal(&cfg); err != nil {
		panic(fmt.Sprintf("Ошибка создания конфига: %s", err.Error()))
	}

	return
}
