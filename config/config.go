package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Schema struct {
	Profiler struct {
		Prometheus    bool   `mapstructure:"prometheus"`
		StatsdAddress string `mapstructure:"statsd_address"`
		Service       string `mapstructure:"service"`
	} `mapstructure:"profiler"`

	Db struct {
		Host     string `mapstructure:"host"`
		User     string `mapstructure:"user"`
		Name     string `mapstructure:"name"`
		Password string `mapstructure:"password"`
		Debug    bool   `mapstructure:"debug"`
	} `mapstructure:"thangtest_database"`

	Paging struct {
		Limit int `mapstructure:"limit"`
	} `mapstructure:"paging"`

	Encryption struct {
		OIDKey           string `mapstructure:"oid_key"`
		JWTSecret        string `mapstructure:"jwt_secret"`
		JWTSecretPartner string `mapstructure:"jwt_secret_partner"`
		JWTExp           int    `mapstructure:"jwt_exp"`
		JWTPol           string `mapstructure:"jwt_pol"`
	} `mapstructure:"encryption"`

	Tenant struct {
		Url          string `mapstructure:"url"`
		ApiKey       string `mapstructure:"api_key"`
		ApiSecretKey string `mapstructure:"api_secret_key"`
	} `mapstructure:"tenant"`
}

var Config Schema

func init() {
	config := viper.New()
	config.SetConfigName("config")
	config.AddConfigPath(".")          // Look for config in current directory
	config.AddConfigPath("config/")    // Optionally look for config in the working directory.
	config.AddConfigPath("../config/") // Look for config needed for tests.
	config.AddConfigPath("../")        // Look for config needed for tests.

	config.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	config.AutomaticEnv()

	err := config.ReadInConfig() // Find and read the config file
	if err != nil {              // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	err = config.Unmarshal(&Config)
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
