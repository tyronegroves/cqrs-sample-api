package config

import (
	"github.com/spf13/viper"
	"strings"
)

type configuration struct {
	LocalAddresses []string `mapstructure:"local_addresses"`
	TrustedProxies []string `mapstructure:"trusted_proxies"`
	EventStoreDb   struct {
		Url string `mapstructure:"url"`
	} `mapstructure:"eventstoredb"`
}

func LoadConfiguration() (*configuration, error) {
	cfg := &configuration{}
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
